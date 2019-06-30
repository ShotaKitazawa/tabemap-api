package sqlite

import (
	"github.com/ShotaKitazawa/tabemap-api/adapter/gateway"
	"github.com/ShotaKitazawa/tabemap-api/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Connect(target string) *gorm.DB {
	var err error

	db, err = gorm.Open("sqlite3", target)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&gateway.Input{}) {
		if err := db.Table("inputs").CreateTable(&gateway.Input{}).Error; err != nil {
			panic(err)
		}
	}

	return db
}

func CloseConn() {
	db.Close()
}

func GetEnv() string {
	return utils.GetEnvOrDefault("SQLITE_FILENAME", "tabemap.sqlite3")
}
