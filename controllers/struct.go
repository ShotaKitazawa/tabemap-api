package controllers

import "time"

type (
	// RequestCreate is request of /article/:title (POST)
	RequestCreate struct {
		Title       string  `json:"title"`
		URL         string  `json:"url"`
		Description string  `json:"description"`
		Type        string  `json:"type"`
		Lat         float64 `json:"latitude"`
		Lng         float64 `json:"longitude"`
	}
	// ResponseCreate is response of /article/:title (POST)
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

	// RequestDetail is request of /article/:id (GET)
	RequestDetail struct {
		ID int64 `json:"id"`
	}
	// ResponseDetail is response of /article/:id (GET)
	ResponseDetail struct {
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

	// RequestSearch is request of /article/search (POST)
	RequestSearch struct {
		ID     int64   `json:"id"`
		Title  string  `json:"title"`
		Type   string  `json:"type"`
		Lat    float64 `json:"latitude"`
		Lng    float64 `json:"longitude"`
		Limit  int     `json:"limit"`
		Offset int     `json:"offset"`
	}
	// ResponseSearch is response of /article/ (GET)
	ResponseSearchOne struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		Type      string    `json:"type"`
		Lat       float64   `json:"latitude"`
		Lng       float64   `json:"longitude"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	ResponseSearch []ResponseSearchOne

	// RequestUpdate is request of /article/:id (PUT)
	RequestUpdate struct {
		ID          int64   `json:"id"`
		Title       string  `json:"title"`
		URL         string  `json:"url"`
		Description string  `json:"description"`
		Type        string  `json:"type"`
		Lat         float64 `json:"latitude"`
		Lng         float64 `json:"longitude"`
	}
	// ResponseUpdate is response of /article/:id (PUT)
	ResponseUpdate struct {
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

	// RequestDelete is request of /article/:id (Delete)
	RequestDelete struct {
		ID int64 `json:"id"`
	}
	ResponseDelete struct {
		ID int64 `json:"id"`
	}
)
