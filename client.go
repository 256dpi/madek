package madek

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

// ErrInvalidAuthentication is returned when the supplied authentication
// credentials have been rejected.
var ErrInvalidAuthentication = errors.New("invalid authentication")

// ErrAccessForbidden is returned when the requested resources is protected.
var ErrAccessForbidden = errors.New("access forbidden")

// ErrRequestFailed is returned when the request failed due to some error.
var ErrRequestFailed = errors.New("request failed")

// ErrNotFound is returned when the requested resource ist not found.
var ErrNotFound = errors.New("not found")

// A Client is used to request data from the Madek API.
type Client struct {
	client       http.Client
	address      string
	username     string
	password     string
	authorCache  map[string]*Author
	groupCache   map[string]*Group
	keywordCache map[string]string
	licenseCache map[string]string
	mutex        sync.Mutex
}

// NewClient will create and return a new Client.
func NewClient(address, username, password string) *Client {
	return &Client{
		address:      address,
		username:     username,
		password:     password,
		authorCache:  make(map[string]*Author),
		groupCache:   make(map[string]*Group),
		keywordCache: make(map[string]string),
		licenseCache: make(map[string]string),
	}
}

// URL appends the passed format to the Madek address.
func (c *Client) URL(format string, args ...interface{}) string {
	args = append([]interface{}{c.address}, args...)
	return fmt.Sprintf("%s"+format, args...)
}

// CompileCollection will fully compile a collection with all available data
// from the API.
func (c *Client) CompileCollection(id string) (*Collection, error) {
	collStr, err := c.Fetch(c.URL("/api/collections/%s", id))
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, gjson.Get(collStr, "created_at").Str)
	if err != nil {
		return nil, err
	}

	coll := &Collection{
		ID:        id,
		CreatedAt: createdAt,
	}

	metaData, err := c.compileMetaData(c.URL("/api/collections/%s/meta-data/", id))
	if err != nil {
		return nil, err
	}

	coll.MetaData = metaData

	var mediaEntryIds []string

	var page = 0
	for {
		mediaEntriesStr, err := c.Fetch(c.URL("/api/media-entries/?collection_id=%s&page=%d", id, page))
		if err != nil {
			return nil, err
		}

		for _, id := range gjson.Get(mediaEntriesStr, "media-entries.#.id").Array() {
			mediaEntryIds = append(mediaEntryIds, id.Str)
		}

		if !gjson.Get(mediaEntriesStr, "_json-roa.collection.next").Exists() {
			break
		}

		page++
	}

	var wg sync.WaitGroup

	asyncErrors := make(chan error, len(mediaEntryIds))
	mediaEntries := make(chan *MediaEntry, len(mediaEntryIds))
	wg.Add(len(mediaEntryIds))

	for _, entryID := range mediaEntryIds {
		go func(id string) {
			defer wg.Done()

			mediaEntry, err := c.CompileMediaEntry(id)
			if err != nil {
				asyncErrors <- err
				return
			}

			mediaEntries <- mediaEntry
		}(entryID)
	}

	wg.Wait()
	close(asyncErrors)
	close(mediaEntries)

	if len(asyncErrors) > 0 {
		return nil, <-asyncErrors
	}

	for mediaEntry := range mediaEntries {
		coll.MediaEntries = append(coll.MediaEntries, mediaEntry)
	}

	return coll, nil
}

// CompileMediaEntry will fully compile a media entry with all available data from the API.
func (c *Client) CompileMediaEntry(id string) (*MediaEntry, error) {
	mediaEntryStr, err := c.Fetch(c.URL("/api/media-entries/%s", id))
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, gjson.Get(mediaEntryStr, "created_at").Str)
	if err != nil {
		return nil, err
	}

	mediaEntry := &MediaEntry{
		ID:        id,
		CreatedAt: createdAt,
	}

	metaData, err := c.compileMetaData(c.URL("/api/media-entries/%s/meta-data/", id))
	if err != nil {
		return nil, err
	}

	mediaEntry.MetaData = metaData

	mediaFileStr, err := c.Fetch(c.URL(gjson.Get(mediaEntryStr, "_json-roa.relations.media-file.href").Str))
	if err != nil {
		return nil, err
	}

	mediaEntry.FileID = gjson.Get(mediaFileStr, "id").Str
	mediaEntry.FileName = gjson.Get(mediaFileStr, "filename").Str
	mediaEntry.StreamURL = c.URL(gjson.Get(mediaFileStr, "_json-roa.relations.data-stream.href").Str)
	mediaEntry.DownloadURL = c.URL("/files/%s", mediaEntry.FileID)

	previewIDs := gjson.Get(mediaFileStr, "previews.#.id").Array()

	var wg sync.WaitGroup

	asyncErrors := make(chan error, len(previewIDs))
	previews := make(chan *Preview, len(previewIDs))
	wg.Add(len(previewIDs))

	for _, previewID := range previewIDs {
		go func(pid string) {
			defer wg.Done()

			previewStr, err := c.Fetch(c.URL("/api/previews/%s", pid))
			if err != nil {
				asyncErrors <- err
				return
			}

			previews <- &Preview{
				ID:          pid,
				Type:        gjson.Get(previewStr, "media_type").Str,
				ContentType: gjson.Get(previewStr, "content_type").Str,
				Size:        gjson.Get(previewStr, "thumbnail").Str,
				Width:       int(gjson.Get(previewStr, "width").Num),
				Height:      int(gjson.Get(previewStr, "height").Num),
				URL:         c.URL("/media/%s", pid),
			}
		}(previewID.Str)
	}

	wg.Wait()
	close(asyncErrors)
	close(previews)

	if len(asyncErrors) > 0 {
		return nil, <-asyncErrors
	}

	for preview := range previews {
		mediaEntry.Previews = append(mediaEntry.Previews, preview)
	}

	return mediaEntry, nil
}

func (c *Client) compileMetaData(url string) (*MetaData, error) {
	metaDataStr, err := c.Fetch(url)
	if err != nil {
		return nil, err
	}

	metaData := &MetaData{}

	for _, metaDatum := range gjson.Get(metaDataStr, "meta-data").Array() {
		metaID := metaDatum.Get("id").Str
		metaKey := metaDatum.Get("meta_key_id").Str

		if !stringInList(supportedMetaKeys, metaKey) {
			continue
		}

		metaDatumStr, err := c.Fetch(c.URL("/api/meta-data/%s", metaID))
		if err != nil {
			return nil, err
		}

		typ := gjson.Get(metaDatumStr, "type").Str

		switch typ {
		case "MetaDatum::Text", "MetaDatum::TextDate":
			strValue := gjson.Get(metaDatumStr, "value").Str

			switch metaKey {
			case "madek_core:title":
				metaData.Title = strValue
			case "madek_core:subtitle":
				metaData.Subtitle = strValue
			case "madek_core:description":
				metaData.Description = strValue
			case "madek_core:portrayed_object_date":
				metaData.Year = strValue
			case "madek_core:copyright_notice":
				metaData.Copyright.Holder = strValue
			case "copyright:copyright_usage":
				metaData.Copyright.Usage = strValue
			default:
				return nil, fmt.Errorf("unhandled meta datum: %s: %s", typ, metaKey)
			}
		case "MetaDatum::Keywords":
			var list []string

			for _, item := range gjson.Get(metaDatumStr, "value.#.id").Array() {
				name, err := c.getKeywordTerm(item.Str)
				if err != nil {
					return nil, err
				}

				list = append(list, name)
			}

			switch metaKey {
			case "madek_core:keywords":
				metaData.Keywords = list
			case "media_content:type":
				metaData.Genres = list
			case "copyright:license":
				metaData.Copyright.Licenses = list
			default:
				return nil, fmt.Errorf("unhandled meta datum: %s: %s", typ, metaKey)
			}
		case "MetaDatum::People":
			switch metaKey {
			case "madek_core:authors":
				authors, err := c.getAuthors(metaDatumStr)
				if err != nil {
					return nil, err
				}

				metaData.Authors = authors
			case "zhdk_bereich:institutional_affiliation":
				groups, err := c.getGroups(metaDatumStr)
				if err != nil {
					return nil, err
				}

				metaData.Affiliation = groups
			default:
				return nil, fmt.Errorf("unhandled meta datum: %s: %s", typ, metaKey)
			}
		default:
			return nil, fmt.Errorf("unhandled meta datum: %s: %s", typ, metaKey)
		}
	}

	return metaData, err
}

func (c *Client) getAuthors(metaDatum string) ([]*Author, error) {
	var authors []*Author

	for _, item := range gjson.Get(metaDatum, "value.#.id").Array() {
		author, err := c.getAuthor(item.Str)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (c *Client) getAuthor(id string) (*Author, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if author, ok := c.authorCache[id]; ok {
		return author, nil
	}

	personStr, err := c.Fetch(c.URL("/api/people/%s", id))
	if err != nil {
		return nil, err
	}

	author := &Author{
		ID:        gjson.Get(personStr, "id").Str,
		FirstName: gjson.Get(personStr, "first_name").Str,
		LastName:  gjson.Get(personStr, "last_name").Str,
	}

	c.authorCache[id] = author

	return author, nil
}

func (c *Client) getGroups(metaDatum string) ([]*Group, error) {
	var authors []*Group

	for _, item := range gjson.Get(metaDatum, "value.#.id").Array() {
		author, err := c.getGroup(item.Str)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (c *Client) getGroup(id string) (*Group, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if group, ok := c.groupCache[id]; ok {
		return group, nil
	}

	groupStr, err := c.Fetch(c.URL("/api/people/%s", id))
	if err != nil {
		return nil, err
	}

	group := &Group{
		ID:        gjson.Get(groupStr, "id").Str,
		Name:      gjson.Get(groupStr, "last_name").Str,
		Pseudonym: gjson.Get(groupStr, "pseudonym").Str,
	}

	c.groupCache[id] = group

	return group, nil
}

func (c *Client) getKeywordTerm(id string) (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if term, ok := c.keywordCache[id]; ok {
		return term, nil
	}

	keywordStr, err := c.Fetch(c.URL("/api/keywords/%s", id))
	if err != nil {
		return "", err
	}

	term := gjson.Get(keywordStr, "term").Str
	c.keywordCache[id] = term

	return term, nil
}

func (c *Client) getLicenseLabel(id string) (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if label, ok := c.licenseCache[id]; ok {
		return label, nil
	}

	licenseStr, err := c.Fetch(c.URL("/api/licenses/%s", id))
	if err != nil {
		return "", err
	}

	label := gjson.Get(licenseStr, "label").Str
	c.licenseCache[id] = label

	return label, nil
}

// Fetch will request the passed URL from Madek.
func (c *Client) Fetch(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Accept", "application/json-roa+json")

	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	switch res.StatusCode {
	case http.StatusUnauthorized:
		return "", ErrInvalidAuthentication
	case http.StatusForbidden:
		return "", ErrAccessForbidden
	case http.StatusNotFound:
		return "", ErrNotFound
	case http.StatusOK:
		return string(bytes), nil
	default:
		return "", ErrRequestFailed
	}
}

func stringInList(list []string, str string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}

	return false
}
