package controllers

import "time"

type (
	RequestCreate struct {
		Title       string  `json:"title"`
		URL         string  `json:"url"`
		Description string  `json:"description"`
		Type        string  `json:"type"`
		Lat         float64 `json:"latitude"`
		Lng         float64 `json:"longitude"`
	}
	ResponseCreate struct {
		ID          int64     `json:"id"`
		Title       string    `json:"title"`
		URL         string    `json:"url"`
		Description string    `json:"description"`
		Type        string    `json:"type"`
		Lat         float64   `json:"latitude"`
		Lng         float64   `json:"longitude"`
		CreatedAt   time.Time `json:"created_at"`
	}

	RequestRead struct {
		ID     int64   `json:"id"`
		Title  string  `json:"title"`
		Type   string  `json:"type"`
		Lat    float64 `json:"latitude"`
		Lng    float64 `json:"longitude"`
		Limit  int     `json:"limit"`
		Offset int     `json:"offset"`
	}
	ResponseRead struct {
		ID          int64     `json:"id"`
		Title       string    `json:"title"`
		URL         string    `json:"url"`
		Description string    `json:"description"`
		Lat         float64   `json:"latitude"`
		Lng         float64   `json:"longitude"`
		Type        string    `json:"type"`
		CreatedAt   time.Time `json:"created_at"`
	}
)
