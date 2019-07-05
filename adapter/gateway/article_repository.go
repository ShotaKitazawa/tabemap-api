package gateway

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/ShotaKitazawa/tabemap-api/domain"
)

type (
	// ArticleRepository is repository
	ArticleRepository struct {
		DBConn *gorm.DB
	}

	// Shop struct = shop table
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

func (r *ArticleRepository) Store(d *domain.Article) (id int64, err error) {
	s := &Shop{
		Name:        d.Title,
		URL:         d.URL,
		Description: d.Description,
		Type:        d.Type,
		Lat:         d.Lat,
		Lng:         d.Lng,
	}
	if errors.New("Name is empty"); s.Name == "" {
		return
	}
	if err = r.DBConn.Create(s).Error; err != nil {
		return
	}

	return int64(s.ID), nil
}
func (r *ArticleRepository) Find(article *domain.Article, limit, offset int) ([]*domain.Article, error) {
	var queryArray []string
	if article.Title != "" {
		queryArray = append(queryArray, fmt.Sprintf("title=\"%s\"", article.Title))
	}
	if article.Type != "" {
		queryArray = append(queryArray, fmt.Sprintf("type=\"%s\"", article.Title))
	}
	if article.Lat != 0 {
		queryArray = append(queryArray, fmt.Sprintf("lat=\"%g\"", article.Lat))
	}
	if article.Lng != 0 {
		queryArray = append(queryArray, fmt.Sprintf("lng=\"%g\"", article.Lng))
	}
	query := strings.Join(queryArray, " and ")

	result := make([]*domain.Article, 0)
	if limit == 0 {
		limit = -1
	}
	r.DBConn.Limit(limit).Offset(offset).Find(result, query)

	return result, nil
}
func (r *ArticleRepository) Update() (*domain.Article, error) {
	return nil, nil
}
func (r *ArticleRepository) Delete() (*domain.Article, error) {
	return nil, nil
}
