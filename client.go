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

// CompileCollection will fully compile a collection with all available data
// from the API.
func (c *Client) CompileCollection(id string) (*Collection, error) {
	// fetch collection
	collStr, err := c.Fetch(c.URL("/api/collections/%s", id))
	if err != nil {
		return nil, err
	}

	// parse time
	createdAt, err := time.Parse(time.RFC3339, gjson.Get(collStr, "created_at").Str)
	if err != nil {
		return nil, err
	}

	// prepare collection
	coll := &Collection{
		ID:        id,
		CreatedAt: createdAt,
	}

	// fetch meta data
	coll.MetaData, err = c.CompileMetaData(c.URL("/api/collections/%s/meta-data/", id))
	if err != nil {
		return nil, err
	}

	// prepare media entries
	var mediaEntryIds []string

	// fetch all media entries
	for page := 0; ; page++ {
		// fetch media entry
		mediaEntriesStr, err := c.Fetch(c.URL("/api/media-entries/?collection_id=%s&page=%d", id, page))
		if err != nil {
			return nil, err
		}

		// append ids
		for _, id := range gjson.Get(mediaEntriesStr, "media-entries.#.id").Array() {
			mediaEntryIds = append(mediaEntryIds, id.Str)
		}

		// check if there is a next page
		if !gjson.Get(mediaEntriesStr, "_json-roa.collection.next").Exists() {
			break
		}
	}

	// prepare wait group
	var wg sync.WaitGroup
	wg.Add(len(mediaEntryIds))

	// prepare result
	asyncErrors := make(chan error, len(mediaEntryIds))
	mediaEntries := make(chan *MediaEntry, len(mediaEntryIds))

	// compile media entries concurrently
	for _, entryID := range mediaEntryIds {
		go func(id string) {
			defer wg.Done()

			// compile media entry
			mediaEntry, err := c.CompileMediaEntry(id)
			if err != nil {
				asyncErrors <- err
				return
			}

			// send media entry
			mediaEntries <- mediaEntry
		}(entryID)
	}

	// await done
	wg.Wait()
	close(asyncErrors)
	close(mediaEntries)

	// check errors
	if len(asyncErrors) > 0 {
		return nil, <-asyncErrors
	}

	// collect media entries
	for mediaEntry := range mediaEntries {
		coll.MediaEntries = append(coll.MediaEntries, mediaEntry)
	}

	return coll, nil
}

// CompileMediaEntry will fully compile a media entry with all available data
// from the API.
func (c *Client) CompileMediaEntry(id string) (*MediaEntry, error) {
	// fetch media entry
	mediaEntryStr, err := c.Fetch(c.URL("/api/media-entries/%s", id))
	if err != nil {
		return nil, err
	}

	// parse time
	createdAt, err := time.Parse(time.RFC3339, gjson.Get(mediaEntryStr, "created_at").Str)
	if err != nil {
		return nil, err
	}

	// prepare media entry
	mediaEntry := &MediaEntry{
		ID:        id,
		CreatedAt: createdAt,
	}

	// compile meta data
	mediaEntry.MetaData, err = c.CompileMetaData(c.URL("/api/media-entries/%s/meta-data/", id))
	if err != nil {
		return nil, err
	}

	// fetch media file
	mediaFileStr, err := c.Fetch(c.URL(gjson.Get(mediaEntryStr, "_json-roa.relations.media-file.href").Str))
	if err != nil {
		return nil, err
	}

	// set file infos
	mediaEntry.FileID = gjson.Get(mediaFileStr, "id").Str
	mediaEntry.FileName = gjson.Get(mediaFileStr, "filename").Str
	mediaEntry.StreamURL = c.URL(gjson.Get(mediaFileStr, "_json-roa.relations.data-stream.href").Str)
	mediaEntry.DownloadURL = c.URL("/files/%s", mediaEntry.FileID)

	// collect previews
	previewIDs := gjson.Get(mediaFileStr, "previews.#.id").Array()

	// prepare wait group
	var wg sync.WaitGroup
	wg.Add(len(previewIDs))

	// prepare result
	asyncErrors := make(chan error, len(previewIDs))
	previews := make(chan *Preview, len(previewIDs))

	// fetch previews concurrently
	for _, previewID := range previewIDs {
		go func(pid string) {
			defer wg.Done()

			// fetch preview
			previewStr, err := c.Fetch(c.URL("/api/previews/%s", pid))
			if err != nil {
				asyncErrors <- err
				return
			}

			// send preview
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

	// await done
	wg.Wait()
	close(asyncErrors)
	close(previews)

	// check errors
	if len(asyncErrors) > 0 {
		return nil, <-asyncErrors
	}

	// collect previews
	for preview := range previews {
		mediaEntry.Previews = append(mediaEntry.Previews, preview)
	}

	return mediaEntry, nil
}

// CompileMetaData will compile the metadata found at the specified url.
func (c *Client) CompileMetaData(url string) (*MetaData, error) {
	// fetch meta data
	metaDataStr, err := c.Fetch(url)
	if err != nil {
		return nil, err
	}

	// prepare meta data
	metaData := &MetaData{}

	// parse all meta datum
	for _, metaDatum := range gjson.Get(metaDataStr, "meta-data").Array() {
		// get id and key
		metaID := metaDatum.Get("id").Str
		metaKey := metaDatum.Get("meta_key_id").Str

		// continue if not supported
		if !stringInList(supportedMetaKeys, metaKey) {
			continue
		}

		// fetch meta datum
		metaDatumStr, err := c.Fetch(c.URL("/api/meta-data/%s", metaID))
		if err != nil {
			return nil, err
		}

		// get type
		typ := gjson.Get(metaDatumStr, "type").Str

		// handle according to type
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
				name, err := c.GetKeywordTerm(item.Str)
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
				// fetch authors
				for _, item := range gjson.Get(metaDatumStr, "value.#.id").Array() {
					author, err := c.GetAuthor(item.Str)
					if err != nil {
						return nil, err
					}
					metaData.Authors = append(metaData.Authors, author)
				}
			case "zhdk_bereich:institutional_affiliation":
				for _, item := range gjson.Get(metaDatumStr, "value.#.id").Array() {
					group, err := c.GetGroup(item.Str)
					if err != nil {
						return nil, err
					}
					metaData.Affiliation = append(metaData.Affiliation, group)
				}
			default:
				return nil, fmt.Errorf("unhandled meta datum: %s: %s", typ, metaKey)
			}
		default:
			return nil, fmt.Errorf("unhandled meta datum: %s: %s", typ, metaKey)
		}
	}

	return metaData, err
}

// GetAuthor will find the author with the provided id.
func (c *Client) GetAuthor(id string) (*Author, error) {
	// acquire mutex
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// check cache
	if author, ok := c.authorCache[id]; ok {
		return author, nil
	}

	// fetch person
	person, err := c.Fetch(c.URL("/api/people/%s", id))
	if err != nil {
		return nil, err
	}

	// prepare author
	author := &Author{
		ID:        gjson.Get(person, "id").Str,
		FirstName: gjson.Get(person, "first_name").Str,
		LastName:  gjson.Get(person, "last_name").Str,
	}

	// cache author
	c.authorCache[id] = author

	return author, nil
}

// GetGroup will find the group with the provided id.
func (c *Client) GetGroup(id string) (*Group, error) {
	// acquire mutex
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// check cache
	if group, ok := c.groupCache[id]; ok {
		return group, nil
	}

	// fetch group
	groupStr, err := c.Fetch(c.URL("/api/people/%s", id))
	if err != nil {
		return nil, err
	}

	// prepare group
	group := &Group{
		ID:        gjson.Get(groupStr, "id").Str,
		Name:      gjson.Get(groupStr, "last_name").Str,
		Pseudonym: gjson.Get(groupStr, "pseudonym").Str,
	}

	// cache group
	c.groupCache[id] = group

	return group, nil
}

// GetKeywordTerm will find the term for the provided keyword id.
func (c *Client) GetKeywordTerm(id string) (string, error) {
	// acquire mutex
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// check cache
	if term, ok := c.keywordCache[id]; ok {
		return term, nil
	}

	// fetch keyword
	keyword, err := c.Fetch(c.URL("/api/keywords/%s", id))
	if err != nil {
		return "", err
	}

	// get term
	term := gjson.Get(keyword, "term").Str
	c.keywordCache[id] = term

	return term, nil
}

// GetLicenseLabel will find the label for the provided license id.
func (c *Client) GetLicenseLabel(id string) (string, error) {
	// acquire mutex
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// check cache
	if label, ok := c.licenseCache[id]; ok {
		return label, nil
	}

	// fetch license
	license, err := c.Fetch(c.URL("/api/licenses/%s", id))
	if err != nil {
		return "", err
	}

	// get label
	label := gjson.Get(license, "label").Str
	c.licenseCache[id] = label

	return label, nil
}

// URL appends the passed format to the Madek address.
func (c *Client) URL(format string, args ...interface{}) string {
	args = append([]interface{}{c.address}, args...)
	return fmt.Sprintf("%s"+format, args...)
}

// Fetch will request the specified URL from Madek.
func (c *Client) Fetch(url string) (string, error) {
	// prepare request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// set headers
	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Accept", "application/json-roa+json")

	// perform request
	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	// ensure body close
	defer res.Body.Close()

	// read body
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// check status code
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
