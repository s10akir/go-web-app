package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() {
	db, err := gorm.Open("sqlite3", "/tmp/go-web-app.sqlite3")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Task{})

	defer db.Close()
}
