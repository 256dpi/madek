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
	Address  string
	Username string
	Password string
}

// NewClient will create and return a new Client.
func NewClient(address, username, password string) *Client {
	return &Client{
		Address:  address,
		Username: username,
		Password: password,
	}
}

// CompileSet will fully compile a set with all available data from the API.
func (c *Client) CompileSet(id string) (*Set, error) {
	setStr, err := c.fetch(c.url("/api/collections/%s", id))
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

	metaDataStr, err := c.fetch(c.url("/api/collections/%s/meta-data/", id))
	if err != nil {
		return nil, err
	}

	metaDataKeys := gjson.Get(metaDataStr, "meta-data.#.meta_key_id").Array()
	metaDataIds := gjson.Get(metaDataStr, "meta-data.#.id").Array()

	for i, key := range metaDataKeys {
		if key.Str == "madek_core:title" {
			metaDatumStr, err := c.fetch(c.url("/api/meta-data/%s", metaDataIds[i].Str))
			if err != nil {
				return nil, err
			}

			set.Title = gjson.Get(metaDatumStr, "value").Str
		}
	}

	var mediaEntryIds []string

	var page = 0
	for {
		mediaEntriesStr, err := c.fetch(c.url("/api/media-entries/?collection_id=%s&page=%d", id, page))
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

			mediaEntry, err := c.compileMediaEntry(id)
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
		set.MediaEntries = append(set.MediaEntries, mediaEntry)
	}

	return set, nil
}

func (c *Client) compileMediaEntry(id string) (*MediaEntry, error) {
	mediaEntryStr, err := c.fetch(c.url("/api/media-entries/%s", id))
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

	metaDataStr, err := c.fetch(c.url("/api/media-entries/%s/meta-data/", id))
	if err != nil {
		return nil, err
	}

	metaDataKeys := gjson.Get(metaDataStr, "meta-data.#.meta_key_id").Array()
	metaDataIds := gjson.Get(metaDataStr, "meta-data.#.id").Array()

	for i, key := range metaDataKeys {
		if key.Str == "madek_core:title" {
			metaDatumStr, err := c.fetch(c.url("/api/meta-data/%s", metaDataIds[i].Str))
			if err != nil {
				return nil, err
			}

			mediaEntry.Title = gjson.Get(metaDatumStr, "value").Str
		}
	}

	mediaFileStr, err := c.fetch(c.url(gjson.Get(mediaEntryStr, "_json-roa.relations.media-file.href").Str))
	if err != nil {
		return nil, err
	}

	mediaEntry.FileID = gjson.Get(mediaFileStr, "id").Str
	mediaEntry.FileName = gjson.Get(mediaFileStr, "filename").Str
	mediaEntry.StreamURL = c.url(gjson.Get(mediaFileStr, "_json-roa.relations.data-stream.href").Str)
	mediaEntry.DownloadURL = c.url("/files/%s", mediaEntry.FileID)

	previewIDs := gjson.Get(mediaFileStr, "previews.#.id").Array()

	for _, previewID := range previewIDs {
		previewStr, err := c.fetch(c.url("/api/previews/%s", previewID.Str))
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
			URL:         c.url("/media/%s", previewID.Str),
		}

		mediaEntry.Previews = append(mediaEntry.Previews, preview)
	}

	return mediaEntry, nil
}

func (c *Client) fetch(path string) (string, error) {
	println(path)

	res, str, err := gorequest.New().Get(path).
		SetBasicAuth(c.Username, c.Password).
		Set("Accept", "application/json-roa+json").
		End()

	if len(err) > 0 {
		return "", &RequestError{
			Err: err[0],
			URL: path,
		}
	}

	if res.StatusCode == http.StatusUnauthorized {
		return "", &RequestError{
			Err: ErrInvalidAuthentication,
			URL: path,
		}
	}

	if res.StatusCode == http.StatusForbidden {
		return "", &RequestError{
			Err: ErrAccessForbidden,
			URL: path,
		}
	}

	if res.StatusCode == http.StatusNotFound {
		return "", &RequestError{
			Err: ErrNotFound,
			URL: path,
		}
	}

	if res.StatusCode != http.StatusOK {
		return "", &RequestError{
			Err: ErrRequestFailed,
			URL: path,
		}
	}

	return str, nil
}

func (c *Client) url(format string, args ...interface{}) string {
	args = append([]interface{}{c.Address}, args...)
	return fmt.Sprintf("%s"+format, args...)
}
