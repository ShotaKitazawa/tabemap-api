package interfaces

import "github.com/ShotaKitazawa/tabemap-api/domain"

type ArticleInteractor interface {
	Add(*domain.Article) (*domain.Article, error)
	Detail(*domain.Article) (*domain.Article, error)
	Find(*domain.Article, int, int) ([]*domain.Article, error)
	Update(*domain.Article) (*domain.Article, error)
	Delete(*domain.Article) (*domain.Article, error)
}
