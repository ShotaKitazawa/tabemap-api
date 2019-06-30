package controllers

import (
	"time"

	"github.com/ShotaKitazawa/tabemap-api/adapter/gateway"
	"github.com/ShotaKitazawa/tabemap-api/adapter/interfaces"
	"github.com/ShotaKitazawa/tabemap-api/usecase"
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
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
		}
		Response struct {
			ID          uint      `json:"id"`
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

	data, err := controller.Interactor.Get()
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "input_controller: cannot get data"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{
		ID:          data.ID,
		Title:       data.Title,
		URL:         data.URL,
		Description: data.Description,
		Lat:         data.Lat,
		Lng:         data.Lng,
		Type:        data.Type,
		CreatedAt:   data.CreatedAt,
	}
	c.JSON(201, res)
}
func (controller *ArticleController) Update(c interfaces.Context) {
}
func (controller *ArticleController) Delete(c interfaces.Context) {
}
