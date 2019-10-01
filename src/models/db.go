package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db, err = gorm.Open("sqlite3", "/tmp/go-web-app.sqlite3")

func Init() {
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Task{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
