package usecase

import (
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/ShotaKitazawa/tabemap-api/usecase/interfaces"
)

type ArticleInteractor struct {
	ArticleRepository interfaces.ArticleRepository
	Logger            interfaces.Logger
}

func (i *ArticleInteractor) Add() (*domain.Article, error) {
	i.Logger.Log("add data")
	return i.ArticleRepository.Store()
}
func (i *ArticleInteractor) Get() (*domain.Article, error) {
	i.Logger.Log("add data")
	return i.ArticleRepository.Store()
}
