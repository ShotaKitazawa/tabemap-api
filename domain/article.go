package domain

import "time"

// Article is Shop data
type Article struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Lat         float64   `json:"latitude"`
	Lng         float64   `json:"longitude"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
