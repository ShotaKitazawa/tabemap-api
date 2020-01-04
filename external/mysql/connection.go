package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	dbrepo "github.com/ShotaKitazawa/tabemap-api/repositories/sqlx"
)

var db *gorm.DB

func Connect(target string) (*gorm.DB, error) {
	var err error

	db, err = gorm.Open("mysql", target)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&dbrepo.Shop{})
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

	return db, nil
}

func CloseConn() {
	db.Close()
}
