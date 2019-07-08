package controllers

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ShotaKitazawa/tabemap-api/adapter/gateway"
	"github.com/ShotaKitazawa/tabemap-api/adapter/interfaces"
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/ShotaKitazawa/tabemap-api/usecase"
)

type ArticleController struct {
	Interactor usecase.ArticleInteractor
}

func NewArticleController(dbConn *gorm.DB, logger interfaces.Logger) *ArticleController {
	return &ArticleController{
		Interactor: usecase.ArticleInteractor{
			ArticleRepository: &gateway.ArticleRepository{
				DBConn: dbConn,
			},
			Logger: logger,
		},
	}
}

func (controller *ArticleController) Create(c interfaces.Context) {
	type (
		Request struct {
			ID          int64   `json:"id"`
			Title       string  `json:"title"`
			URL         string  `json:"url"`
			Description string  `json:"description"`
			Lat         float64 `json:"latitude"`
			Lng         float64 `json:"longitude"`
			Type        string  `json:"type"`
		}
		Response struct {
			ID          int64     `json:"id"`
			Title       string    `json:"title"`
			URL         string    `json:"url"`
			Description string    `json:"description"`
			Lat         float64   `json:"latitude"`
			Lng         float64   `json:"longitude"`
			Type        string    `json:"type"`
			CreatedAt   time.Time `json:"created_at"`
		}
	)
	req := Request{}
	c.Bind(&req)

	article := &domain.Article{
		ID:    req.ID,
		Title: req.Title,
		Lat:   req.Lat,
		Lng:   req.Lng,
		Type:  req.Type,
	}
	result, err := controller.Interactor.Add(article)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "input_controller: cannot get data"))
		c.JSON(500, NewError(500, err.Error()))
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
	}
	c.JSON(201, res)

	return
}

func (controller *ArticleController) Read(c interfaces.Context) {
	type (
		Request struct {
			ID     int64   `json:"id"`
			Title  string  `json:"title"`
			Lat    float64 `json:"latitude"`
			Lng    float64 `json:"longitude"`
			Type   string  `json:"type"`
			Limit  int     `json:"limit"`
			Offset int     `json:"offset"`
		}
		Response struct {
			ID          int64     `json:"id"`
			Title       string    `json:"title"`
			URL         string    `json:"url"`
			Description string    `json:"description"`
			Lat         float64   `json:"latitude"`
			Lng         float64   `json:"longitude"`
			Type        string    `json:"type"`
			CreatedAt   time.Time `json:"created_at"`
		}
	)
	req := Request{}
	c.Bind(&req)

	article := &domain.Article{
		ID:    req.ID,
		Title: req.Title,
		Lat:   req.Lat,
		Lng:   req.Lng,
		Type:  req.Type,
	}
	results, err := controller.Interactor.Get(article, req.Limit, req.Offset)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "input_controller: cannot get data"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}

	res := []Response{}
	for _, result := range results {
		res = append(res, Response{
			ID:          result.ID,
			Title:       result.Title,
			URL:         result.URL,
			Description: result.Description,
			Lat:         result.Lat,
			Lng:         result.Lng,
			Type:        result.Type,
			CreatedAt:   result.CreatedAt,
		})
	}
	c.JSON(201, res)

	return
}

func (controller *ArticleController) Update(c interfaces.Context) {
	type (
		Request struct {
			ID          int64   `json:"id"`
			Title       string  `json:"title"`
			URL         string  `json:"url"`
			Description string  `json:"description"`
			Lat         float64 `json:"latitude"`
			Lng         float64 `json:"longitude"`
			Type        string  `json:"type"`
		}
		Response struct {
			ID          int64     `json:"id"`
			Title       string    `json:"title"`
			URL         string    `json:"url"`
			Description string    `json:"description"`
			Lat         float64   `json:"latitude"`
			Lng         float64   `json:"longitude"`
			Type        string    `json:"type"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
		}
	)
	req := Request{}
	c.Bind(&req)

	article := &domain.Article{
		ID:    req.ID,
		Title: req.Title,
		Lat:   req.Lat,
		Lng:   req.Lng,
		Type:  req.Type,
	}
	result, err := controller.Interactor.Update(article)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "input_controller: cannot get data"))
		c.JSON(500, NewError(500, err.Error()))
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
	c.JSON(201, res)

	return
}

func (controller *ArticleController) Delete(c interfaces.Context) {
	return
}
