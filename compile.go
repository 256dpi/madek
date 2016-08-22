package madek

import (
	"time"

	"github.com/tidwall/gjson"
)

func (c *Client) CompileSet(id string) (*Set, error) {
	setStr, err := c.fetch(c.url("collections/%s", id))
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, gjson.Get(setStr, "created_at").Value().(string))
	if err != nil {
		return nil, err
	}

	set := &Set{
		ID:        id,
		CreatedAt: createdAt,
	}

	metaDataStr, err := c.fetch(c.url("collections/%s/meta-data/", id))
	if err != nil {
		return nil, err
	}

	metaDataKeys := gjson.Get(metaDataStr, "meta-data.#.meta_key_id").Multi
	metaDataIds := gjson.Get(metaDataStr, "meta-data.#.id").Multi

	for i, key := range metaDataKeys {
		if key.String() == "madek_core:title" {
			metaDatumStr, err := c.fetch(c.url("meta-data/%s", metaDataIds[i].String()))
			if err != nil {
				return nil, err
			}

			set.Name = gjson.Get(metaDatumStr, "value").String()
		}
	}

	mediaEntryIds := make([]string, 0)

	var page = 0
	for {
		mediaEntriesStr, err := c.fetch(c.url("media-entries/?collection_id=%s&page=%d", id, page))
		if err != nil {
			return nil, err
		}

		for _, id := range gjson.Get(mediaEntriesStr, "media-entries.#.id").Value().([]interface{}) {
			mediaEntryIds = append(mediaEntryIds, id.(string))
		}

		if !gjson.Get(mediaEntriesStr, "_json-roa.collection.next").Exists() {
			break
		}

		page++
	}

	for _, entryID := range mediaEntryIds {
		mediaEntryStr, err := c.fetch(c.url("media-entries/%s", entryID))
		if err != nil {
			return nil, err
		}

		createdAt, err := time.Parse(time.RFC3339, gjson.Get(mediaEntryStr, "created_at").Value().(string))
		if err != nil {
			return nil, err
		}

		media := Media{
			ID:        entryID,
			CreatedAt: createdAt,
		}

		metaDataStr, err := c.fetch(c.url("media-entries/%s/meta-data/", entryID))
		if err != nil {
			return nil, err
		}

		metaDataKeys := gjson.Get(metaDataStr, "meta-data.#.meta_key_id").Multi
		metaDataIds := gjson.Get(metaDataStr, "meta-data.#.id").Multi

		for i, key := range metaDataKeys {
			if key.String() == "madek_core:title" {
				metaDatumStr, err := c.fetch(c.url("meta-data/%s", metaDataIds[i].String()))
				if err != nil {
					return nil, err
				}

				media.Name = gjson.Get(metaDatumStr, "value").String()
			}
		}

		set.Media = append(set.Media, media)
	}

	return set, nil
}
