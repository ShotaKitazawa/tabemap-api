package gateway

import (
	"database/sql/driver"
	"os"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestArticleController(t *testing.T) {
	t.Run("Store()", func(t *testing.T) {
		t.Run("保存する(正常系)", func(t *testing.T) {
			var r ArticleRepository
			var d *domain.Article
			db, mock, err := getDBMock()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			db.LogMode(true)

			r = ArticleRepository{DBConn: db}
			d = &domain.Article{
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

			result, err := r.Store(d)
			assert.Nil(t, err)
			if result != nil {
				result.CreatedAt = d.CreatedAt
				result.UpdatedAt = d.UpdatedAt
			}
			assert.Equal(t, d, result)
		})
		t.Run("保存する(異常系)", func(t *testing.T) {
			var r ArticleRepository
			var d *domain.Article
			db, _, err := getDBMock()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			db.LogMode(true)

			r = ArticleRepository{DBConn: db}
			d = &domain.Article{
				ID:          1,
				Title:       "",
				URL:         "example.com",
				Description: "fot test",
				Type:        "Japanese",
				Lat:         1.1,
				Lng:         1.1,
			}

			result, err := r.Store(d)
			assert.NotNil(t, err)
			assert.Equal(t, (*domain.Article)(nil), result)
		})
	})
	t.Run("Find()", func(t *testing.T) {
		t.Run("すべて取得する", func(t *testing.T) {
			var r ArticleRepository
			var d *domain.Article
			db, mock, err := getDBMock()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			db.LogMode(true)

			r = ArticleRepository{DBConn: db}
			d = &domain.Article{
				ID:    0,
				Title: "hoge",
				Lat:   1,
				Lng:   1,
				Type:  "",
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `shops`")).
				WillReturnRows(sqlmock.NewRows([]string{"id", "title", "url", "description", "type", "lat", "lng"}).AddRow(d.ID, d.Title, d.URL, d.Description, d.Type, d.Lat, d.Lng))

			_, err = r.Search(d, 0, 0)
			assert.Nil(t, err)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
		t.Run("name=`ほげ`を取得する", func(t *testing.T) {
			var r ArticleRepository
			var d *domain.Article
			db, mock, err := getDBMock()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			db.LogMode(true)

			r = ArticleRepository{DBConn: db}
			d = &domain.Article{
				ID:    0,
				Title: "ほげ",
				Lat:   0,
				Lng:   0,
				Type:  "",
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `shops` WHERE `shops`.`deleted_at` IS NULL AND ((name=\"ほげ\"))")).
				WillReturnRows(sqlmock.NewRows([]string{"id", "title", "url", "description", "type", "lat", "lng"}).AddRow(d.ID, d.Title, d.URL, d.Description, d.Type, d.Lat, d.Lng))

			_, err = r.Search(d, 0, 0)
			assert.Nil(t, err)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
	})
	t.Run("Update()", func(t *testing.T) {
		/*
			t.Run("ID=1のnameを更新する", func(t *testing.T) {
				var r ArticleRepository
				var d *domain.Article
				db, mock, err := getDBMock()
				if err != nil {
					t.Fatal(err)
				}
				defer db.Close()
				db.LogMode(true)

				r = ArticleRepository{DBConn: db}
				d = &domain.Article{
					ID:    1,
					Title: "ほげ",
					Lat:   0,
					Lng:   0,
					Type:  "",
				}

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `shops` SET `id` = ?, `name` = ?, `updated_at` = ?")).
					WithArgs(d.ID, d.Title, d.UpdatedAt, d.ID).
					WillReturnResult(sqlmock.NewResult(d.ID, 1))
				mock.ExpectCommit()

				_, err = r.Update(d)
				assert.Nil(t, err)
				// TODO: occured error
				//assert.Nil(t, mock.ExpectationsWereMet())
			})
		*/
	})
	t.Run("Delete()", func(t *testing.T) {
		// TODO
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

	//gdb, err := gorm.Open("sqlite3", db)
	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}
