package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func DbInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	return db
}

func Init() {
	Client = DbInit()
}
