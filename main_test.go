package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"github.com/ShotaKitazawa/tabemap-api/external"
	_ "github.com/ShotaKitazawa/tabemap-api/external"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	if err := exec.Command("sqlite3", "../../inputs-test.sqlite3").Run(); err != nil {
		panic(err)
	}
	r = external.Router
	code := m.Run()
	os.Exit(code)
}

func TestArticleController(t *testing.T) {
	t.Run("Create->Create->ReadAll->ReadOne->Update->Delete->ReadAll", func(t *testing.T) {
		var req *http.Request
		var rec *httptest.ResponseRecorder

		req = httptest.NewRequest("POST", "/api/article", nil)
		rec = httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "hogefuga", rec.Body.String())

		req = httptest.NewRequest("POST", "/api/article", nil)
		rec = httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "hogefuga", rec.Body.String())

		req = httptest.NewRequest("GET", "/api/article", nil)
		rec = httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "hogefuga", rec.Body.String())
	})
}
