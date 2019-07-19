// +build integration

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"github.com/ShotaKitazawa/tabemap-api/domain"
	"github.com/ShotaKitazawa/tabemap-api/external"
)

var r *gin.Engine

// test on sqlite3
func TestMain(m *testing.M) {
	r = external.Router
	code := m.Run()
	os.Exit(code)
}

func TestArticleController(t *testing.T) {
	t.Run("Create->Create->ReadAll->ReadOne->Update->Delete->ReadAll", func(t *testing.T) {
		var req *http.Request
		var rec *httptest.ResponseRecorder
		var b []byte
		var request, expectedResponse, actualResponse domain.Article
		// var reqDomainArticles, resDomainArticles, expecteds []domain.Article
		var err error

		// 1. Create
		request = domain.Article{
			Title:       "test",
			URL:         "http://example.com",
			Description: "for test",
			Type:        "Japanese",
			Lat:         1.1,
			Lng:         -1.1,
		}
		expectedResponse = domain.Article{
			ID:          1,
			Title:       "test",
			URL:         "http://example.com",
			Description: "for test",
			Type:        "Japanese",
			Lat:         1.1,
			Lng:         -1.1,
		}
		b, err = json.Marshal(&request)
		assert.Nil(t, err)
		req = httptest.NewRequest("POST", "/api/article", bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)
		json.NewDecoder(rec.Body).Decode(&actualResponse)
		actualResponse.CreatedAt = expectedResponse.CreatedAt
		actualResponse.UpdatedAt = expectedResponse.UpdatedAt
		assert.Equal(t, expectedResponse, actualResponse)

		// 2. Create
		request = domain.Article{
			Title:       "test",
			URL:         "http://example.net",
			Description: "for test",
			Type:        "Chinese",
			Lat:         1.1,
			Lng:         -1.1,
		}
		expectedResponse = domain.Article{
			ID:          2,
			Title:       "test",
			URL:         "http://example.net",
			Description: "for test",
			Type:        "Chinese",
			Lat:         1.1,
			Lng:         -1.1,
		}
		b, err = json.Marshal(&request)
		assert.Nil(t, err)
		req = httptest.NewRequest("POST", "/api/article", bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)
		json.NewDecoder(rec.Body).Decode(&actualResponse)
		actualResponse.CreatedAt = expectedResponse.CreatedAt
		actualResponse.UpdatedAt = expectedResponse.UpdatedAt
		assert.Equal(t, expectedResponse, actualResponse)

		// 3. ReadAll
		// 4. ReadOne
		// 5. Update
		// 6. Delete
		// 7. ReadAll
	})
}

/*
% curl localhost:8080/api/article -X PUT -H "Content-Type: application/json" -d '{"id": 1, "title": "update", "url": "http://example.net", "description": "hoge", "latitude": -1, "longitude": 3, "type": "chinese"}'
% curl localhost:8080/api/article -X DELETE -H "Content-Type: application/json" -d '{"id": 1}'
% curl localhost:8080/api/article -X POST -H "Content-Type: application/json" -d '{"title": "test", "url": "http://example.com", "description": "for test", "latitude": 1, "longitude": 2, "type": "japanese"}'
% curl localhost:8080/api/article/
*/

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
