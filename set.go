package madek

import "time"

type Set struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Media     []Media   `json:"media"`
}

type Media struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	FileID      string    `json:"file_id"`
	StreamURL   string    `json:"stream_url"`
	DownloadURL string    `json:"file_url"`
	Previews    []Preview `json:"previews"`
}

type Preview struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	ContentType string `json:"content_type"`
	Size        string `json:"size"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	URL         string `json:"url"`
}
