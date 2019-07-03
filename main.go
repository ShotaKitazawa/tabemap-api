package main

import (
	"github.com/ShotaKitazawa/tabemap-api/external"
	// "github.com/ShotaKitazawa/tabemap-api/external/mysql"
	"github.com/ShotaKitazawa/tabemap-api/external/sqlite"
)

func main() {
	defer sqlite.CloseConn()
	// defer mysql.CloseConn()

	external.Router.Run("0.0.0.0:8080")
}
