package gateway

import (
	"database/sql/driver"
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

type AnyTime struct{} // I don't actually know if I even need this

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("sqlite3", db)
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
			d := &domain.Article{
				ID:          1,
				Title:       "hoge",
				URL:         "example.com",
				Description: "fot test",
				Type:        "japanese",
				Lat:         1.1,
				Lng:         1.1,
			}
			var shop_id, pos_id int64
			shop_id = 1
			pos_id = 1

			mock.ExpectBegin()
			mock.ExpectExec("^INSERT INTO \"shops\" (.+)$").
				WithArgs(AnyTime{}, AnyTime{}, nil, d.Title, d.URL, d.Description, d.Type).
				WillReturnResult(sqlmock.NewResult(shop_id, 1))
			mock.ExpectCommit()
			mock.ExpectBegin()
			mock.ExpectExec("^INSERT INTO \"positions\" (.+)$").
				WithArgs(AnyTime{}, AnyTime{}, nil, d.Lat, d.Lng).
				WillReturnResult(sqlmock.NewResult(pos_id, 1))
			mock.ExpectCommit()
			mock.ExpectBegin()
			mock.ExpectExec("^INSERT INTO \"maps\" (.+)$").
				WithArgs(AnyTime{}, AnyTime{}, nil, shop_id, pos_id).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			_, err = r.Store(d)
			assert.Nil(t, err)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	})
}
