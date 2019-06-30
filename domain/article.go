package domain

import "time"

// Article is Shop data
type Article struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Lat         float64   `json:"latitude"`
	Lng         float64   `json:"longitude"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}
