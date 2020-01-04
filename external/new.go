package external

import (
	"github.com/jinzhu/gorm"

	"github.com/ShotaKitazawa/tabemap-api/controllers"
	dbrepo "github.com/ShotaKitazawa/tabemap-api/repositories/sqlx"
	"github.com/ShotaKitazawa/tabemap-api/usecase"
	"github.com/ShotaKitazawa/tabemap-api/usecase/interfaces"
)

func NewArticleController(dbConn *gorm.DB, logger interfaces.Logger) *controllers.ArticleController {
	return &controllers.ArticleController{
		Interactor: usecase.ArticleInteractor{
			DBRepository: &dbrepo.Repository{
				DBConn: dbConn,
				Logger: logger,
			},
			Logger: logger,
		},
	}
}
