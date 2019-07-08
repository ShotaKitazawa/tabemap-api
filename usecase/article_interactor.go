package usecase

import (
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/ShotaKitazawa/tabemap-api/usecase/interfaces"
)

type ArticleInteractor struct {
	ArticleRepository interfaces.ArticleRepository
	Logger            interfaces.Logger
}

func (i *ArticleInteractor) Add(d *domain.Article) (*domain.Article, error) {
	i.Logger.Log("add data")
	return i.ArticleRepository.Store(d)
}
func (i *ArticleInteractor) Get(article *domain.Article, limit, offset int) ([]*domain.Article, error) {
	i.Logger.Log("store data")
	return i.ArticleRepository.Find(article, limit, offset)
}

func (i *ArticleInteractor) Update(article *domain.Article) (*domain.Article, error) {
	i.Logger.Log("update data")
	return i.ArticleRepository.Update(article)
}
