package controllers

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/ShotaKitazawa/tabemap-api/controllers/interfaces"
	"github.com/ShotaKitazawa/tabemap-api/domain"
)

type ArticleController struct {
	Interactor interfaces.ArticleInteractor
	Logger     interfaces.Logger
}

func (controller *ArticleController) Create(c interfaces.Context) {
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
	controller.Logger.Debug(req)
	controller.Logger.Debug(article)

	result, err := controller.Interactor.Add(article)
	if err != nil {
		controller.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
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

func (controller *ArticleController) Detail(c interfaces.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		controller.Logger.Error(errors.Wrap(err, "article_controller: cannot cast string to int64"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	article := &domain.Article{
		ID: id,
	}
	result, err := controller.Interactor.Detail(article)
	if err != nil {
		controller.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := ResponseDetail{
		ID:          result.ID,
		Title:       result.Title,
		URL:         result.Type,
		Description: result.Description,
		Type:        result.Type,
		Lat:         result.Lat,
		Lng:         result.Lng,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}
	c.JSON(http.StatusOK, res)

	return
}

func (controller *ArticleController) Search(c interfaces.Context) {
	req := RequestSearch{}
	c.Bind(&req)

	article := &domain.Article{
		ID:    req.ID,
		Title: req.Title,
		Type:  req.Type,
		Lat:   req.Lat,
		Lng:   req.Lng,
	}
	controller.Logger.Debug(req)
	controller.Logger.Debug(article)

	results, err := controller.Interactor.Find(article, req.Limit, req.Offset)
	if err != nil {
		controller.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := ResponseSearch{}
	for _, result := range results {
		res = append(res, ResponseSearchOne{
			ID:        result.ID,
			Title:     result.Title,
			Type:      result.Type,
			Lat:       result.Lat,
			Lng:       result.Lng,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, res)

	return
}

func (controller *ArticleController) Update(c interfaces.Context) {
	req := RequestUpdate{}
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
	controller.Logger.Debug(req)
	controller.Logger.Debug(article)

	result, err := controller.Interactor.Update(article)
	if err != nil {
		controller.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	res := ResponseUpdate{
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
	req := RequestDelete{}
	c.Bind(&req)

	article := &domain.Article{
		ID: req.ID,
	}
	controller.Logger.Debug(req)
	controller.Logger.Debug(article)

	_, err := controller.Interactor.Delete(article)
	if err != nil {
		controller.Logger.Error(errors.Wrap(err, "article_controller: cannot get data"))
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusAccepted, nil)

	return
}
