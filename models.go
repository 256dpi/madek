package madek

import "time"

// A Set contains multiple media entries.
type Set struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	CreatedAt    time.Time     `json:"created_at"`
	MediaEntries []*MediaEntry `json:"media_entries"`
}

// A MediaEntry contains multiple previews.
type MediaEntry struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
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
