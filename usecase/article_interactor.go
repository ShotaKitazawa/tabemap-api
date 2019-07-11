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
	i.Logger.Info("add data")
	return i.ArticleRepository.Store(d)
}

func (i *ArticleInteractor) Find(article *domain.Article, limit, offset int) ([]*domain.Article, error) {
	i.Logger.Info("get data")
	return i.ArticleRepository.Search(article, limit, offset)
}

func (i *ArticleInteractor) Update(article *domain.Article) (*domain.Article, error) {
	i.Logger.Info("update data")
	return i.ArticleRepository.Update(article)
}

func (i *ArticleInteractor) Delete(article *domain.Article) (*domain.Article, error) {
	i.Logger.Info("delete data")
	return i.ArticleRepository.Delete(article)
}
