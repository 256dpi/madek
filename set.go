package madek

import "time"

type Set struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Media     []Media   `json:"media"`
}

type Media struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
