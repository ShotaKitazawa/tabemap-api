package interfaces

import "github.com/ShotaKitazawa/tabemap-api/domain"

type ArticleRepository interface {
	Store(*domain.Article) (int64, error)
	Find(*domain.Article, int, int) ([]*domain.Article, error)
	Update(*domain.Article) (error)
	Delete() (*domain.Article, error)
}
