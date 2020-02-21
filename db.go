package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"minx.com/demo/app"
)

// InitDB
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("DB OPEN ERROR!")
	}

	db.Debug()
	db.AutoMigrate(&app.App{})

	return db
}
