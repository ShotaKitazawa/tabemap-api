package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ShotaKitazawa/tabemap-api/adapter/gateway"
	"github.com/ShotaKitazawa/tabemap-api/utils"
)

var db *gorm.DB

func Connect(target string) *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", target)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&gateway.Map{}) {
		if err := db.Table("map").CreateTable(&gateway.Map{}).Error; err != nil {
			fmt.Println(err)
		}
	}
	if !db.HasTable(&gateway.Shop{}) {
		if err := db.Table("shop").CreateTable(&gateway.Shop{}).Error; err != nil {
			fmt.Println(err)
		}
	}
	if !db.HasTable(&gateway.Position{}) {
		if err := db.Table("position").CreateTable(&gateway.Position{}).Error; err != nil {
			fmt.Println(err)
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
