package interfaces

import "github.com/ShotaKitazawa/tabemap-api/domain"

type ArticleRepository interface {
	Store(*domain.Article) (*domain.Article, error)
	Find(*domain.Article, int, int) ([]*domain.Article, error)
	Update(*domain.Article) (*domain.Article,error)
	Delete(int) (*domain.Article,error)
}
