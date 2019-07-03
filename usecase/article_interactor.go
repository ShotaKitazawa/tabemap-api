package usecase

import (
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/ShotaKitazawa/tabemap-api/usecase/interfaces"
)

type ArticleInteractor struct {
	ArticleRepository interfaces.ArticleRepository
	Logger            interfaces.Logger
}

func (i *ArticleInteractor) Add(d *domain.Article) (uint64, error) {
	i.Logger.Log("add data")
	return i.ArticleRepository.Store(d)
}
func (i *ArticleInteractor) Get() (*domain.Article, error) {
	i.Logger.Log("add data")
	return i.ArticleRepository.FindByName()
}
