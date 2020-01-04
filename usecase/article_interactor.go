package usecase

import (
	"fmt"

	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/ShotaKitazawa/tabemap-api/usecase/interfaces"
)

type ArticleInteractor struct {
	DBRepository interfaces.DBRepository
	Logger       interfaces.Logger
}

func (i *ArticleInteractor) Add(d *domain.Article) (*domain.Article, error) {
	i.Logger.Info("add data")
	return i.DBRepository.Store(d)
}

func (i *ArticleInteractor) Detail(article *domain.Article) (*domain.Article, error) {
	i.Logger.Info("get data")
	result, err := i.DBRepository.Search(article, 0, 0)
	if err != nil {
		return nil, err
	} else if len(result) != 1 {
		return nil, fmt.Errorf("")
	}
	return result[0], nil
}

func (i *ArticleInteractor) Find(article *domain.Article, limit, offset int) ([]*domain.Article, error) {
	i.Logger.Info("get data")
	return i.DBRepository.Search(article, limit, offset)
}

func (i *ArticleInteractor) Update(article *domain.Article) (*domain.Article, error) {
	i.Logger.Info("update data")
	return i.DBRepository.Update(article)
}

func (i *ArticleInteractor) Delete(article *domain.Article) (*domain.Article, error) {
	i.Logger.Info("delete data")
	return i.DBRepository.Delete(article)
}
