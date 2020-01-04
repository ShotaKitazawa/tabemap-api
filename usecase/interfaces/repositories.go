package interfaces

import "github.com/ShotaKitazawa/tabemap-api/domain"

type DBRepository interface {
	Store(*domain.Article) (*domain.Article, error)
	Search(*domain.Article, int, int) ([]*domain.Article, error)
	Update(*domain.Article) (*domain.Article, error)
	Delete(*domain.Article) (*domain.Article, error)
}
