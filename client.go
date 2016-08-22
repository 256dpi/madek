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

var ErrInvalidAuthentication = errors.New("Invalid Authentication")
var ErrAccessForbidden = errors.New("Access Forbidden")
var ErrRequestFailed = errors.New("Request Failed")
var ErrNotFound = errors.New("Not Found")

type Client struct {
	Address  string
	Username string
	Password string
}

func NewClient(address, username, password string) *Client {
	return &Client{
		Address:  address,
		Username: username,
		Password: password,
	}
}

func (c *Client) CompileSet(id string, loadMediaEntries bool) (*Set, error) {
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

	metaDataKeys := gjson.Get(metaDataStr, "meta-data.#.meta_key_id").Multi
	metaDataIds := gjson.Get(metaDataStr, "meta-data.#.id").Multi

	for i, key := range metaDataKeys {
		if key.Str == "madek_core:title" {
			metaDatumStr, err := c.fetch(c.url("/api/meta-data/%s", metaDataIds[i].Str))
			if err != nil {
				return nil, err
			}

			set.Title = gjson.Get(metaDatumStr, "value").Str
		}
	}

	mediaEntryIds := make([]string, 0)

	var page = 0
	for {
		mediaEntriesStr, err := c.fetch(c.url("/api/media-entries/?collection_id=%s&page=%d", id, page))
		if err != nil {
			return nil, err
		}

		for _, id := range gjson.Get(mediaEntriesStr, "media-entries.#.id").Multi {
			mediaEntryIds = append(mediaEntryIds, id.Str)
		}

		if !gjson.Get(mediaEntriesStr, "_json-roa.collection.next").Exists() {
			break
		}

		page++
	}

	if loadMediaEntries {
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
			set.MediaEntries = append(set.MediaEntries, mediaEntry)
		}
	}

	return set, nil
}

func (c *Client) CompileMediaEntry(id string) (*MediaEntry, error) {
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

	metaDataKeys := gjson.Get(metaDataStr, "meta-data.#.meta_key_id").Multi
	metaDataIds := gjson.Get(metaDataStr, "meta-data.#.id").Multi

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

	previewIDs := gjson.Get(mediaFileStr, "previews.#.id").Multi

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
		return "", err[0]
	}

	if res.StatusCode == http.StatusUnauthorized {
		return "", ErrInvalidAuthentication
	}

	if res.StatusCode == http.StatusForbidden {
		return "", ErrAccessForbidden
	}

	if res.StatusCode == http.StatusNotFound {
		return "", ErrNotFound
	}

	if res.StatusCode != http.StatusOK {
		return "", ErrRequestFailed
	}

	return str, nil
}

func (c *Client) url(format string, args ...interface{}) string {
	args = append([]interface{}{c.Address}, args...)
	return fmt.Sprintf("%s"+format, args...)
}
