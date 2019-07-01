package mysql

import (
	"github.com/ShotaKitazawa/tabemap-api/adapter/gateway"
	"github.com/ShotaKitazawa/tabemap-api/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect(target string) *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", target)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&gateway.Article{}) {
		if err := db.Table("article").CreateTable(&gateway.Article{}).Error; err != nil {
			panic(err)
		}
	}

	return db
}

func CloseConn() {
	db.Close()
}

func GetEnv() string {
	dbUser := utils.GetEnvOrDefault("DB_USER", "root")
	dbPass := utils.GetEnvOrDefault("DB_PASSWORD", "password")
	dbHost := utils.GetEnvOrDefault("DB_HOST", "localhost")
	dbPort := utils.GetEnvOrDefault("DB_PORT", "3306")
	dbName := utils.GetEnvOrDefault("DB_NAME", "tabemap")
	return dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}
