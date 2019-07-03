package gateway

import (
	"os"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestArticleController(t *testing.T) {
	t.Run("Store()", func(t *testing.T) {
		t.Run("xxを保存する", func(t *testing.T) {
			db, mock, err := getDBMock()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			db.LogMode(true)

			r := ArticleRepository{DBConn: db}
			var map_id, shop_id, pos_id uint
			var name, url, description, shop_type string
			var lat, lng float64

			mock.ExpectBegin()
			// insert into map
			map_id = 1
			shop_id = 1
			pos_id = 1
			mock.ExpectQuery(regexp.QuoteMeta(
				`INSERT INTO "map" ("id","shop_id","pos_id") VALUES ($1,$2, $3)`)).
				WithArgs(map_id, shop_id, pos_id).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(map_id))
			// insert into shop
			shop_id = 1
			name = "hoge"
			url = "example.com"
			description = "for test"
			shop_type = "japanese"
			mock.ExpectQuery(regexp.QuoteMeta(
				`INSERT INTO "shop" ("id","name", "url", "description", "type") VALUES ($1,$2,$3,$4,$5)`)).
				WithArgs(shop_id, name, url, description, shop_type).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(shop_id))
			// insert into position
			pos_id = 1
			lat = 1
			lng = 1
			mock.ExpectQuery(regexp.QuoteMeta(
				`INSERT INTO "pos" ("id","lat","lng") VALUES ($1,$2, $3)`)).
				WithArgs(pos_id, lat, lng).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(pos_id))
			mock.ExpectCommit()

			_, err = r.Store()
			if err != nil {
				t.Fatal(err)
			}
		})
	})
}
