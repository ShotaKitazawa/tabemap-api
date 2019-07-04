package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/ShotaKitazawa/tabemap-api/adapter/gateway"
	"github.com/ShotaKitazawa/tabemap-api/utils"
)

var db *gorm.DB

func Connect(target string) *gorm.DB {
	var err error

	db, err = gorm.Open("sqlite3", target)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&gateway.Shop{})
	/*
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
	*/

	return db
}

func CloseConn() {
	db.Close()
}

func GetEnv() string {
	return utils.GetEnvOrDefault("SQLITE_FILENAME", "tabemap.sqlite3")
}
