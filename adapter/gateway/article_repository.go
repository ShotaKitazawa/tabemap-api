package gateway

import (
	"github.com/ShotaKitazawa/tabemap-api/domain"

	"github.com/jinzhu/gorm"
)

type (
	ArticleRepository struct {
		DBConn *gorm.DB
	}

	Article struct {
		gorm.Model
		Title       string
		URL         string
		Description string
		Lat         float64
		Lng         float64
		Type        string
	}
)

func (r *ArticleRepository) Store() (*domain.Article, error) {
	return nil, nil
}
func (r *ArticleRepository) FindByName() (*domain.Article, error) {
	return nil, nil
}
func (r *ArticleRepository) FindByType() (*domain.Article, error) {
	return nil, nil
}
func (r *ArticleRepository) FindByLngLat() (*domain.Article, error) {
	return nil, nil
}
func (r *ArticleRepository) FindByLocate() (*domain.Article, error) {
	return nil, nil
}
func (r *ArticleRepository) Update() (*domain.Article, error) {
	return nil, nil
}
func (r *ArticleRepository) Delete() (*domain.Article, error) {
	return nil, nil
}
