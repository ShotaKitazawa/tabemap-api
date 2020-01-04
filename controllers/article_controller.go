package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/ShotaKitazawa/tabemap-api/controllers/interfaces"
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/ShotaKitazawa/tabemap-api/usecase"
)

type ArticleController struct {
	Interactor usecase.ArticleInteractor
}

// TODO 用途のぶんだけハンドラを用意する
// c.f.
//  現在: 検索 -> Read()
//  TODO: 検索(名前) -> SearchName(), 検索(ID) -> SearchID()

func (controller *ArticleController) CreateArticle(c interfaces.Context) {
	req := RequestCreate{}
	c.Bind(&req)

	article := &domain.Article{
		Title:       req.Title,
		URL:         req.URL,
		Description: req.Description,
		Type:        req.Type,
		Lat:         req.Lat,
		Lng:         req.Lng,
	}
	controller.Interactor.Logger.Debug(req)
	controller.Interactor.Logger.Debug(article)

	result, err := controller.Interactor.Add(article)
	if err != nil {
		controller.Interactor.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := ResponseCreate{
		ID:          result.ID,
		Title:       result.Title,
		URL:         result.URL,
		Description: result.Description,
		Type:        result.Type,
		Lat:         result.Lat,
		Lng:         result.Lng,
		CreatedAt:   result.CreatedAt,
	}
	c.JSON(http.StatusCreated, res)

	return
}

//func (controller *ArticleController) Detail(c interfaces.Context) {
//}

func (controller *ArticleController) Read(c interfaces.Context) {
	type (
	)
	req := RequestRead{}
	c.Bind(&req)

	article := &domain.Article{
		ID:    req.ID,
		Title: req.Title,
		Type:  req.Type,
		Lat:   req.Lat,
		Lng:   req.Lng,
	}
	controller.Interactor.Logger.Debug(req)
	controller.Interactor.Logger.Debug(article)

	if idStr := c.Param("id"); idStr != "" {
		if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
			controller.Interactor.Logger.Error(errors.Wrap(err, "article_controller: cannot cast string to int64"))
			c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
			return
		} else {
			article.ID = id
		}
	}
	results, err := controller.Interactor.Find(article, req.Limit, req.Offset)
	if err != nil {
		controller.Interactor.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := []Response{}
	for _, result := range results {
		res = append(res, Response{
			ID:          result.ID,
			Title:       result.Title,
			URL:         result.URL,
			Description: result.Description,
			Type:        result.Type,
			Lat:         result.Lat,
			Lng:         result.Lng,
			CreatedAt:   result.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, res)

	return
}

func (controller *ArticleController) Update(c interfaces.Context) {
	type (
		Request struct {
			ID          int64   `json:"id"`
			Title       string  `json:"title"`
			URL         string  `json:"url"`
			Description string  `json:"description"`
			Type        string  `json:"type"`
			Lat         float64 `json:"latitude"`
			Lng         float64 `json:"longitude"`
		}
		Response struct {
			ID          int64     `json:"id"`
			Title       string    `json:"title"`
			URL         string    `json:"url"`
			Description string    `json:"description"`
			Type        string    `json:"type"`
			Lat         float64   `json:"latitude"`
			Lng         float64   `json:"longitude"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
		}
	)
	req := Request{}
	c.Bind(&req)

	article := &domain.Article{
		ID:          req.ID,
		Title:       req.Title,
		URL:         req.URL,
		Description: req.Description,
		Type:        req.Type,
		Lat:         req.Lat,
		Lng:         req.Lng,
	}
	controller.Interactor.Logger.Debug(req)
	controller.Interactor.Logger.Debug(article)

	result, err := controller.Interactor.Update(article)
	if err != nil {
		controller.Interactor.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := Response{
		ID:          result.ID,
		Title:       result.Title,
		URL:         result.URL,
		Description: result.Description,
		Lat:         result.Lat,
		Lng:         result.Lng,
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}
	c.JSON(http.StatusAccepted, res)

	return
}

func (controller *ArticleController) Delete(c interfaces.Context) {
	type (
		Request struct {
			ID int64 `json:"id"`
		}
		Response struct {
			ID          int64   `json:"id"`
			Title       string  `json:"title"`
			URL         string  `json:"url"`
			Description string  `json:"description"`
			Type        string  `json:"type"`
			Lat         float64 `json:"latitude"`
			Lng         float64 `json:"longitude"`
		}
	)
	req := Request{}
	c.Bind(&req)

	article := &domain.Article{
		ID: req.ID,
	}
	controller.Interactor.Logger.Debug(req)
	controller.Interactor.Logger.Debug(article)

	result, err := controller.Interactor.Delete(article)
	if err != nil {
		controller.Interactor.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := Response{
		ID:          result.ID,
		Title:       result.Title,
		URL:         result.URL,
		Description: result.Description,
		Type:        result.Type,
		Lat:         result.Lat,
		Lng:         result.Lng,
	}
	c.JSON(http.StatusAccepted, res)

	return
}
