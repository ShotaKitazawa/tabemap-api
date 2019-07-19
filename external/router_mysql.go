// +build mysql
// +build !sqlite

package external

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ShotaKitazawa/tabemap-api/adapter/controllers"
	"github.com/ShotaKitazawa/tabemap-api/external/mysql"
	"github.com/ShotaKitazawa/tabemap-api/utils"
)

// Router called by main.go
var Router *gin.Engine

func init() {
	r := gin.Default()
	r.Use(cors.Default())

	v := r.Group("/api")

	logger := &Logger{}
	logger.setLoglevel(utils.GetEnvOrDefault("LOG_MIN_LEVEL", "info"))

	dbConn := mysql.Connect(mysql.GetEnv())

	ArticleController := controllers.NewArticleController(dbConn, logger)

	v.POST("/article", func(c *gin.Context) { ArticleController.Create(c) })
	v.GET("/article", func(c *gin.Context) { ArticleController.Read(c) })
	v.GET("/article/:id", func(c *gin.Context) { ArticleController.Read(c) })
	v.PUT("/article", func(c *gin.Context) { ArticleController.Update(c) })
	v.DELETE("/article", func(c *gin.Context) { ArticleController.Delete(c) })

	Router = r
}

func Run(addr string) {
	defer mysql.CloseConn()
	Router.Run(addr)
}
