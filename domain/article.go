package domain

import "time"

// Article is Shop data
type Article struct {
	ID          int64
	Title       string
	URL         string
	Description string
	Type        string
	Lat         float64
	Lng         float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
