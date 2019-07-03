package gateway

import (
	"github.com/jinzhu/gorm"

	"github.com/ShotaKitazawa/tabemap-api/domain"
)

type (
	// ArticleRepository is repository
	ArticleRepository struct {
		DBConn *gorm.DB
	}

	// Map struct = map table
	Map struct {
		gorm.Model
		ShopID uint
		PosID  uint
	}
	// Shop struct = shop table
	Shop struct {
		gorm.Model
		Name        string
		URL         string
		Description string
		Type        string
	}
	// Position struct = position table
	Position struct {
		gorm.Model
		Lat float64
		Lng float64
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
