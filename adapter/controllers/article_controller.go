package controllers

import (
	"fmt"
	"net/http"
	"strconv"
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
		}
	)
	req := Request{}
	c.Bind(&req)

	article := &domain.Article{
		Title:       req.Title,
		URL:         req.URL,
		Description: req.Description,
		Type:        req.Type,
		Lat:         req.Lat,
		Lng:         req.Lng,
	}
	fmt.Println(req)
	fmt.Println(article)

	result, err := controller.Interactor.Add(article)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "article_controller: cannot get data"))
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
		CreatedAt:   result.CreatedAt,
	}
	c.JSON(http.StatusCreated, res)

	return
}

//func (controller *ArticleController) Detail(c interfaces.Context) {
//}

func (controller *ArticleController) Read(c interfaces.Context) {
	type (
		Request struct {
			ID     int64   `json:"id"`
			Title  string  `json:"title"`
			Type   string  `json:"type"`
			Lat    float64 `json:"latitude"`
			Lng    float64 `json:"longitude"`
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
		Type:  req.Type,
		Lat:   req.Lat,
		Lng:   req.Lng,
	}
	fmt.Println(req)
	fmt.Println(article)

	if idStr := c.Param("id"); idStr != "" {
		if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
			controller.Interactor.Logger.Log(errors.Wrap(err, "article_controller: cannot cast string to int64"))
			c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
			return
		} else {
			article.ID = id
		}
	}
	results, err := controller.Interactor.Find(article, req.Limit, req.Offset)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "article_controller: cannot get data"))
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
	fmt.Println(req)
	fmt.Println(article)

	result, err := controller.Interactor.Update(article)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "article_controller: cannot get data"))
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
	fmt.Println(req)
	fmt.Println(article)

	result, err := controller.Interactor.Delete(article)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "article_controller: cannot get data"))
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
