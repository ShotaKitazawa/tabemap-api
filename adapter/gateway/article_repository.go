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
		ShopID uint64
		PosID  uint64
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

func (r *ArticleRepository) Store(d *domain.Article) (id uint64, err error) {
	shop := &Shop{
		Name:        d.Title,
		URL:         d.URL,
		Description: d.Description,
		Type:        d.Type,
	}
	if err = r.DBConn.Create(shop).Error; err != nil {
		return
	}
	pos := Position{
		Lat: d.Lat,
		Lng: d.Lng,
	}
	if err = r.DBConn.Create(pos).Error; err != nil {
		return
	}
	m := &Map{
		ShopID: uint64(shop.ID),
		PosID:  uint64(pos.ID),
	}
	if err = r.DBConn.Create(m).Error; err != nil {
		return
	}

	return uint64(m.ID), nil
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
