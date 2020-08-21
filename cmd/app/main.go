package main

import (
	"app/internal/database"
	"app/internal/routes"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize routes and begin server
	database.ConnectionString = "./app.db"
	var err error
	database.DB, err = sql.Open("sqlite3", database.ConnectionString)

	defer database.DB.Close()

	if err != nil {
		panic(err.Error())
	}
	routes.InitServer(":8000")
}
