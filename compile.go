package madek

import (
	"bytes"
	"time"

	"encoding/json"
	"github.com/kr/pretty"
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

	metaDataStr, err := c.fetch(c.url("collections/%s/meta-data/?meta_keys[0]=madek_core:title", id))
	if err != nil {
		return nil, err
	}

	metaDatumStr, err := c.fetch(c.url("meta-data/%s", gjson.Get(metaDataStr, "meta-data.0.id").String()))
	if err != nil {
		return nil, err
	}

	set.Name = gjson.Get(metaDatumStr, "value").String()

	buf := new(bytes.Buffer)
	json.Indent(buf, []byte(metaDatumStr), "", "  ")
	pretty.Println(buf.String())

	return set, nil
}
