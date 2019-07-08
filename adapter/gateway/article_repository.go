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

func (r *ArticleRepository) Store(article *domain.Article) (result *domain.Article, err error) {
	shop := &Shop{
		Name:        article.Title,
		URL:         article.URL,
		Description: article.Description,
		Type:        article.Type,
		Lat:         article.Lat,
		Lng:         article.Lng,
	}
	if err = ErrNameIsEmpty; shop.Name == "" {
		return
	}
	if err = r.DBConn.Create(shop).Error; err != nil {
		return
	}
	return &domain.Article{
		ID:          int64(shop.ID),
		Title:       shop.Name,
		URL:         shop.URL,
		Description: shop.Description,
		Type:        shop.Type,
		Lat:         shop.Lat,
		Lng:         shop.Lng,
		CreatedAt:   shop.CreatedAt,
	}, nil
}

func (r *ArticleRepository) Find(article *domain.Article, limit, offset int) (results []*domain.Article, err error) {
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
	if err = r.DBConn.Limit(limit).Offset(offset).Find(&shops, query).Error; err != nil {
		return
	}
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
			CreatedAt:   val.CreatedAt,
		})
	}
	return d, nil
}

func (r *ArticleRepository) Update(article *domain.Article) (result *domain.Article, err error) {
	shop := &Shop{
		Name:        article.Title,
		URL:         article.URL,
		Description: article.Description,
		Type:        article.Type,
		Lat:         article.Lat,
		Lng:         article.Lng,
	}
	query := make(map[string]interface{})
	if article.Title != "" {
		query["name"] = article.Title
	}
	if article.Type != "" {
		query["type"] = article.Type
	}
	if article.Lat != 0 {
		query["lat"] = article.Lat
	}
	if article.Lng != 0 {
		query["lng"] = article.Lng
	}
	if err = r.DBConn.Model(&shop).Updates(query).Error; err != nil {
		return
	}
	return &domain.Article{
		ID:          int64(shop.ID),
		Title:       shop.Name,
		URL:         shop.URL,
		Description: shop.Description,
		Type:        shop.Type,
		Lat:         shop.Lat,
		Lng:         shop.Lng,
		CreatedAt:   shop.CreatedAt,
		UpdatedAt:   shop.UpdatedAt,
	}, nil
}

func (r *ArticleRepository) Delete(id int) (result *domain.Article, err error) {
	return nil, nil
}
