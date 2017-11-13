package madek

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
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

// ErrUnhandledMetaDatum is returned when a received meta datum has not been
// handled.
var ErrUnhandledMetaDatum = errors.New("unhandled meta datum")

// ErrUnhandledMetaDatumType is returned when a fetched meta datum has not
// been handled.
var ErrUnhandledMetaDatumType = errors.New("unhandled meta datum type")

// A Client is used to request data from the Madek API.
type Client struct {
	Address     string
	Username    string
	Password    string
	LogRequests bool

	peopleCache  map[string]string
	keywordCache map[string]string
	licenseCache map[string]string

	shortLock sync.Mutex
}

// NewClient will create and return a new Client.
func NewClient(address, username, password string) *Client {
	return &Client{
		Address:      address,
		Username:     username,
		Password:     password,
		peopleCache:  make(map[string]string),
		keywordCache: make(map[string]string),
		licenseCache: make(map[string]string),
	}
}

// URL appends the passed format to the Madek address.
func (c *Client) URL(format string, args ...interface{}) string {
	args = append([]interface{}{c.Address}, args...)
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
				return nil, errors.Wrap(ErrUnhandledMetaDatum, metaKey)
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
			default:
				return nil, errors.Wrap(ErrUnhandledMetaDatum, metaKey)
			}
		case "MetaDatum::People":
			var list []string

			for _, item := range gjson.Get(metaDatumStr, "value.#.id").Array() {
				name, err := c.getNameOfPerson(item.Str)
				if err != nil {
					return nil, err
				}

				list = append(list, name)
			}

			switch metaKey {
			case "madek_core:authors":
				metaData.Authors = list
			default:
				return nil, errors.Wrap(ErrUnhandledMetaDatum, metaKey)
			}
		case "MetaDatum::Licenses":
			var list []string

			for _, item := range gjson.Get(metaDatumStr, "value.#.id").Array() {
				name, err := c.getLicenseLabel(item.Str)
				if err != nil {
					return nil, err
				}

				list = append(list, name)
			}

			switch metaKey {
			case "copyright:license":
				metaData.Copyright.Licenses = list
			default:
				return nil, errors.Wrap(ErrUnhandledMetaDatum, metaKey)
			}
		default:
			return nil, errors.Wrap(ErrUnhandledMetaDatumType, typ)
		}
	}

	return metaData, err
}

func (c *Client) getNameOfPerson(id string) (string, error) {
	c.shortLock.Lock()
	defer c.shortLock.Unlock()

	if name, ok := c.peopleCache[id]; ok {
		return name, nil
	}

	personStr, err := c.Fetch(c.URL("/api/people/%s", id))
	if err != nil {
		return "", err
	}

	name := gjson.Get(personStr, "first_name").Str + " " + gjson.Get(personStr, "last_name").Str
	c.peopleCache[id] = name

	return name, nil
}

func (c *Client) getKeywordTerm(id string) (string, error) {
	c.shortLock.Lock()
	defer c.shortLock.Unlock()

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
	c.shortLock.Lock()
	defer c.shortLock.Unlock()

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
	if c.LogRequests {
		println("Fetching: " + url)
	}

	res, str, err := gorequest.New().Get(url).
		SetBasicAuth(c.Username, c.Password).
		Set("Accept", "application/json-roa+json").
		End()

	if len(err) > 0 {
		return "", errors.Wrap(err[0], url)
	}

	switch res.StatusCode {
	case http.StatusUnauthorized:
		return "", errors.Wrap(ErrInvalidAuthentication, url)
	case http.StatusForbidden:
		return "", errors.Wrap(ErrAccessForbidden, url)
	case http.StatusNotFound:
		return "", errors.Wrap(ErrNotFound, url)
	case http.StatusOK:
		return str, nil
	default:
		return "", errors.Wrap(ErrRequestFailed, url)
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
