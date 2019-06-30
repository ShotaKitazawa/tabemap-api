package interfaces

import "github.com/ShotaKitazawa/tabemap-api/domain"

type ArticleRepository interface {
	Store() (*domain.Article, error)
	FindByName() (*domain.Article, error)
	FindByType() (*domain.Article, error)
	FindByLngLat() (*domain.Article, error)
	FindByLocate() (*domain.Article, error)
	Update() (*domain.Article, error)
	Delete() (*domain.Article, error)
}
