package gateway

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/ShotaKitazawa/tabemap-api/domain"
)

var ErrNameIsEmpty error

func init() {
	ErrNameIsEmpty = errors.New("Name is empty")
}

type (
	// ArticleRepository is repository
	ArticleRepository struct {
		DBConn *gorm.DB
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

func (r *ArticleRepository) Store(d *domain.Article) (id int64, err error) {
	s := &Shop{
		Name:        d.Title,
		URL:         d.URL,
		Description: d.Description,
		Type:        d.Type,
		Lat:         d.Lat,
		Lng:         d.Lng,
	}
	if err = ErrNameIsEmpty; s.Name == "" {
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
		queryArray = append(queryArray, fmt.Sprintf("name=\"%s\"", article.Title))
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

	shops := []Shop{}
	if limit == 0 {
		limit = -1
	}

	r.DBConn.Limit(limit).Offset(offset).Find(&shops, query)

	var d []*domain.Article
	for _, val := range shops {
		d = append(d, &domain.Article{
			ID:          int64(val.ID),
			Title:       val.Name,
			URL:         val.URL,
			Description: val.Description,
			Type:        val.Type,
			Lat:         val.Lat,
			Lng:         val.Lng,
		})
	}

	return d, nil
}
func (r *ArticleRepository) Update(article *domain.Article) error {
	updates := make(map[string]interface{})
	if article.Title != "" {
		updates["name"] = article.Title
	}
	if article.Type != "" {
		updates["type"] = article.Type
	}
	if article.Lat != 0 {
		updates["lat"] = article.Lat
	}
	if article.Lng != 0 {
		updates["lng"] = article.Lng
	}

	r.DBConn.Model(&article).Updates(updates)
	return nil
}
func (r *ArticleRepository) Delete() (*domain.Article, error) {
	return nil, nil
}
