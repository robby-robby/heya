package main

import (
	"context"
	"database/sql"
	"heya/db"
	"heya/lgg"
)

// Rename to App?
type App struct {
	db *db.Queries
}

func NewApp(db *db.Queries) *App {
	return &App{db: db}
}

var DefaultSettings = db.CreateSettingsParams{
	Codify: false,
	Model:  "gpt-4",
	Editor: "nvim",
	Temp:   10,
}

func (a *App) Bootstrap() error {
	if !a.Exists() {
		if err := a.db.CreateSettings(context.Background(), DefaultSettings); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Exists() bool {
	_, err := a.db.GetSettings(context.Background())
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		lgg.Panic(err)
	}
	return true

}
