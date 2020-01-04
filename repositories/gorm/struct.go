package dbrepo

import (
	"github.com/ShotaKitazawa/tabemap-api/repositories/interfaces"
	"github.com/jinzhu/gorm"
)

type (
	Repository struct {
		DBConn *gorm.DB
		Logger interfaces.Logger
	}

	// Shop struct is DB shop table
	Shop struct {
		gorm.Model
		Name        string
		URL         string
		Description string
		Type        string
		Lat         float64
		Lng         float64
	}
)
