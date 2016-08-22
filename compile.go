package madek

import (
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

func (c *Client) CompileSet(id string, loadMediaEntries bool) (*Set, error) {
	setStr, err := c.fetch(c.url("/api/collections/%s", id))
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, gjson.Get(setStr, "created_at").String())
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
		if key.String() == "madek_core:title" {
			metaDatumStr, err := c.fetch(c.url("/api/meta-data/%s", metaDataIds[i].String()))
			if err != nil {
				return nil, err
			}

			set.Title = gjson.Get(metaDatumStr, "value").String()
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

		errors := make(chan error, len(mediaEntryIds))
		mediaEntries := make(chan *MediaEntry, len(mediaEntryIds))

		wg.Add(len(mediaEntryIds))

		for _, entryID := range mediaEntryIds {
			go func() {
				defer wg.Done()

				mediaEntry, err := c.CompileMediaEntry(entryID)
				if err != nil {
					errors <- err
					return
				}

				mediaEntries <- mediaEntry
			}()
		}

		wg.Wait()
		close(errors)
		close(mediaEntries)

		if len(errors) > 0 {
			return nil, <-errors
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

	createdAt, err := time.Parse(time.RFC3339, gjson.Get(mediaEntryStr, "created_at").String())
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
		if key.String() == "madek_core:title" {
			metaDatumStr, err := c.fetch(c.url("/api/meta-data/%s", metaDataIds[i].String()))
			if err != nil {
				return nil, err
			}

			mediaEntry.Title = gjson.Get(metaDatumStr, "value").String()
		}
	}

	mediaFileStr, err := c.fetch(c.url(gjson.Get(mediaEntryStr, "_json-roa.relations.media-file.href").String()))
	if err != nil {
		return nil, err
	}

	mediaEntry.FileID = gjson.Get(mediaFileStr, "id").String()
	mediaEntry.StreamURL = c.url(gjson.Get(mediaFileStr, "_json-roa.relations.data-stream.href").String())
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
