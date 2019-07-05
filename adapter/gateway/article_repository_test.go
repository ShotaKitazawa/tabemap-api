package gateway

import (
	"database/sql/driver"
	"errors"
	"os"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestArticleController(t *testing.T) {
	t.Run("Store()", func(t *testing.T) {
		t.Run("保存する(正常系)", func(t *testing.T) {
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
				Type:        "Japanese",
				Lat:         1.1,
				Lng:         1.1,
			}

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `shops` (`created_at`,`updated_at`,`deleted_at`,`name`,`url`,`description`,`type`,`lat`,`lng`) VALUES (?,?,?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, d.Title, d.URL, d.Description, d.Type, d.Lat, d.Lng).
				WillReturnResult(sqlmock.NewResult(d.ID, 1))
			mock.ExpectCommit()

			_, err = r.Store(d)
			assert.Nil(t, err)
		})
		t.Run("保存する(異常系)", func(t *testing.T) {
			db, mock, err := getDBMock()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			db.LogMode(true)

			r := ArticleRepository{DBConn: db}
			d := &domain.Article{
				ID:          1,
				Title:       "",
				URL:         "example.com",
				Description: "fot test",
				Type:        "Japanese",
				Lat:         1.1,
				Lng:         1.1,
			}

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `shops` (`created_at`,`updated_at`,`deleted_at`,`name`,`url`,`description`,`type`,`lat`,`lng`) VALUES (?,?,?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, d.Title, d.URL, d.Description, d.Type, d.Lat, d.Lng).
				WillReturnError(errors.New("Name is empty"))
			mock.ExpectCommit()

			_, err = r.Store(d)
			assert.Nil(t, err)
		})
	})
	t.Run("Find()", func(t *testing.T) {
		t.Run("すべて取得する", func(t *testing.T) {
		})
	})
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

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
