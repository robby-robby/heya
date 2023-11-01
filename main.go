package main

import (
	"heya/config"
	"heya/db"
	"heya/lgg"
	"heya/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dbc, err := sqlite.NewOpenDB(config.Config.Dsn)
	if err != nil {
		lgg.Panicf("Failed to open database: %v", err)
	} else {
		defer dbc.Close()
	}

	db := db.New(dbc)
	err = NewApp(db).Bootstrap()
	if err != nil {
		lgg.Panicf("Failed to bootstrap app: %v", err)
	}
}
