package madek

import "time"

var supportedMetaKeys = []string{
	"madek_core:title",
	"madek_core:subtitle",
	"madek_core:description",
	"madek_core:authors",
	"madek_core:keywords",
	"media_content:type",
	"madek_core:portrayed_object_date",
	"madek_core:copyright_notice",
	"copyright:copyright_usage",
	"copyright:license",
	"zhdk_bereich:institutional_affiliation",
}

// Author contains info about an author.
type Author struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// Group contains info about a group.
type Group struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Pseudonym string `json:"pseudonym,omitempty"`
}

// Copyright contains copyright infos.
type Copyright struct {
	Holder   string   `json:"holder,omitempty"`
	Usage    string   `json:"usage,omitempty"`
	Licenses []string `json:"licenses,omitempty"`
}

// MetaData contains multiple metadata key value pairs.
type MetaData struct {
	Title       string    `json:"title,omitempty"`
	Subtitle    string    `json:"subtitle,omitempty"`
	Description string    `json:"description,omitempty"`
	Authors     []*Author `json:"authors,omitempty"`
	Keywords    []string  `json:"keywords,omitempty"`
	Genres      []string  `json:"genres,omitempty"`
	Year        string    `json:"year,omitempty"`
	Copyright   Copyright `json:"copyright,omitempty"`
	Affiliation []*Group  `json:"affiliation,omitempty"`
}

// A Collection contains multiple media entries.
type Collection struct {
	ID           string        `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	MetaData     *MetaData     `json:"meta_data"`
	MediaEntries []*MediaEntry `json:"media_entries"`
}

// A MediaEntry contains multiple previews.
type MediaEntry struct {
	ID          string     `json:"id"`
	MetaData    *MetaData  `json:"meta_data"`
	CreatedAt   time.Time  `json:"created_at"`
	FileID      string     `json:"file_id"`
	FileName    string     `json:"file_name"`
	StreamURL   string     `json:"stream_url"`
	DownloadURL string     `json:"download_url"`
	Previews    []*Preview `json:"previews"`
}

// A Preview is the final accessible media.
type Preview struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	ContentType string `json:"content_type"`
	Size        string `json:"size"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	URL         string `json:"url"`
}
