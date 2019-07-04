package usecase

import (
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/ShotaKitazawa/tabemap-api/usecase/interfaces"
)

type ArticleInteractor struct {
	ArticleRepository interfaces.ArticleRepository
	Logger            interfaces.Logger
}

func (i *ArticleInteractor) Add(d *domain.Article) (int64, error) {
	i.Logger.Log("add data")
	return i.ArticleRepository.Store(d)
}
func (i *ArticleInteractor) Get(article *domain.Article, start, end int) ([]*domain.Article, error) {
	i.Logger.Log("add data")
	return i.ArticleRepository.Find(article, start, end)
}
