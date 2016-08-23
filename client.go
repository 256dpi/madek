package madek

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

// ErrInvalidAuthentication is returned when the supplied authentication
// credentials have been rejected.
var ErrInvalidAuthentication = errors.New("Invalid Authentication")

// ErrAccessForbidden is returned when the requested resources is protected.
var ErrAccessForbidden = errors.New("Access Forbidden")

// ErrRequestFailed is returned when the request failed due to some error.
var ErrRequestFailed = errors.New("Request Failed")

// ErrNotFound ist returned when the requested resource ist not found.
var ErrNotFound = errors.New("Not Found")

// A RequestError ist returned on a request error and provides additional info.
type RequestError struct {
	Err error
	URL string
}

func (r *RequestError) Error() string {
	return r.Err.Error() + ": " + r.URL
}

// A Client is used to request data from the Madek API.
type Client struct {
	Address     string
	Username    string
	Password    string
	LogRequests bool
	MetaKeys    map[string]string
}

// NewClient will create and return a new Client.
func NewClient(address, username, password string) *Client {
	return &Client{
		Address:  address,
		Username: username,
		Password: password,
		MetaKeys: map[string]string{
			"madek_core:title":    "title",
			"madek_core:subtitle": "subtitle",
			//"media_content:type": "type",
			//"zhdk_bereich:institutional_affiliation": "affiliation",
		},
	}
}

// URL appends the passed format to the Madek address.
func (c *Client) URL(format string, args ...interface{}) string {
	args = append([]interface{}{c.Address}, args...)
	return fmt.Sprintf("%s"+format, args...)
}

// CompileSet will fully compile a set with all available data from the API.
func (c *Client) CompileSet(id string) (*Set, error) {
	setStr, err := c.Fetch(c.URL("/api/collections/%s", id))
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, gjson.Get(setStr, "created_at").Str)
	if err != nil {
		return nil, err
	}

	set := &Set{
		ID:        id,
		CreatedAt: createdAt,
	}

	metaData, err := c.compileMetaData(c.URL("/api/collections/%s/meta-data/", id))
	if err != nil {
		return nil, err
	}

	set.MetaData = metaData

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
		set.MediaEntries = append(set.MediaEntries, *mediaEntry)
	}

	return set, nil
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

	for _, previewID := range previewIDs {
		previewStr, err := c.Fetch(c.URL("/api/previews/%s", previewID.Str))
		if err != nil {
			return nil, err
		}

		preview := &Preview{
			ID:          previewID.Str,
			Type:        gjson.Get(previewStr, "media_type").Str,
			ContentType: gjson.Get(previewStr, "content_type").Str,
			Size:        gjson.Get(previewStr, "thumbnail").Str,
			Width:       int(gjson.Get(previewStr, "width").Num),
			Height:      int(gjson.Get(previewStr, "height").Num),
			URL:         c.URL("/media/%s", previewID.Str),
		}

		mediaEntry.Previews = append(mediaEntry.Previews, *preview)
	}

	return mediaEntry, nil
}

func (c *Client) compileMetaData(url string) (MetaData, error) {
	metaDataStr, err := c.Fetch(url)
	if err != nil {
		return nil, err
	}

	metaDataKeys := gjson.Get(metaDataStr, "meta-data.#.meta_key_id").Array()
	metaDataIds := gjson.Get(metaDataStr, "meta-data.#.id").Array()

	metaData := make(MetaData)

	for i, rawKey := range metaDataKeys {
		for madekKey, mapKey := range c.MetaKeys {
			if rawKey.Str == madekKey {
				metaDatumStr, err := c.Fetch(c.URL("/api/meta-data/%s", metaDataIds[i].Str))
				if err != nil {
					return nil, err
				}

				println(metaDatumStr)

				metaData[mapKey] = gjson.Get(metaDatumStr, "value").Str
			}
		}
	}

	return metaData, err
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
		return "", &RequestError{
			Err: err[0],
			URL: url,
		}
	}

	if res.StatusCode == http.StatusUnauthorized {
		return "", &RequestError{
			Err: ErrInvalidAuthentication,
			URL: url,
		}
	}

	if res.StatusCode == http.StatusForbidden {
		return "", &RequestError{
			Err: ErrAccessForbidden,
			URL: url,
		}
	}

	if res.StatusCode == http.StatusNotFound {
		return "", &RequestError{
			Err: ErrNotFound,
			URL: url,
		}
	}

	if res.StatusCode != http.StatusOK {
		return "", &RequestError{
			Err: ErrRequestFailed,
			URL: url,
		}
	}

	return str, nil
}
