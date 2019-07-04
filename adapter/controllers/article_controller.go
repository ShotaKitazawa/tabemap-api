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
}

func (controller *ArticleController) Read(c interfaces.Context) {
	type (
		Request struct {
			ID    int64   `json:"id"`
			Title string  `json:"title"`
			Lat   float64 `json:"latitude"`
			Lng   float64 `json:"longitude"`
			Type  string  `json:"type"`
			Start int     `json:"start"`
			End   int     `json:"end"`
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

	data, err := controller.Interactor.Get(article, req.Start, req.End)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "input_controller: cannot get data"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := []Response{}
	for _, val := range data {
		res = append(res, Response{
			ID:          val.ID,
			Title:       val.Title,
			URL:         val.URL,
			Description: val.Description,
			Lat:         val.Lat,
			Lng:         val.Lng,
			Type:        val.Type,
			CreatedAt:   val.CreatedAt,
		})
	}
	c.JSON(201, res)
}
func (controller *ArticleController) Update(c interfaces.Context) {
}
func (controller *ArticleController) Delete(c interfaces.Context) {
}
