package external

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/ShotaKitazawa/tabemap-api/external/mysql"
	"github.com/ShotaKitazawa/tabemap-api/external/sqlite"
)

// Router called by main.go
var Router *gin.Engine

// Run is called by main.go (entrypoint)
func Run(ctx context.Context) {
	// Get Value From cmd
	dbType, err := getContextString(ctx, "db-type")
	if err != nil {
		panic(err)
	}
	dsn, err := getContextString(ctx, "dsn")
	if err != nil {
		panic(err)
	}
	bindHost, err := getContextString(ctx, "bind-host")
	if err != nil {
		panic(err)
	}

	// Create DB connection
	var dbConn *gorm.DB
	switch dbType {
	case "sqlite3":
		dbConn, err = sqlite.Connect(dsn)
	case "mysql":
		dbConn, err = mysql.Connect(dsn)
	default:
		panic(fmt.Errorf("No such DB type"))
	}
	if err != nil {
		panic(err)
	}

	// Create Logger
	logger := &Logger{}

	// Create Controllers
	ArticleController := NewArticleController(dbConn, logger)

	// Create *gin.Router
	r := gin.New()
	r.Use(gin.Recovery(), Log(), cors.Default())
	v := r.Group("/api")

	// define handlers
	v.POST("/article", func(c *gin.Context) { ArticleController.Create(c) })
	v.GET("/article", func(c *gin.Context) { ArticleController.Read(c) })
	v.GET("/article/:id", func(c *gin.Context) { ArticleController.Read(c) })
	v.PUT("/article", func(c *gin.Context) { ArticleController.Update(c) })
	v.DELETE("/article", func(c *gin.Context) { ArticleController.Delete(c) })

	// bind
	r.Run(bindHost)
}

func getContextString(ctx context.Context, key string) (result string, err error) {
	inter := ctx.Value(key)
	if inter == nil {
		err = errors.New(fmt.Sprintf("context not in value \"%s\"", key))
		return
	}
	result, ok := inter.(string)
	if !ok {
		err = errors.New(fmt.Sprintf("value \"%s\" is not string", key))
		return
	}
	return
}
