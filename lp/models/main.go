package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabase() *gorm.DB {
	if db == nil {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

		if err != nil {
			panic("Failed to connect to database!")
		}

		db.AutoMigrate(&Agent{})

	}
	return db
}
