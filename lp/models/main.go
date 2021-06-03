package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

		if err != nil {
			panic("Failed to connect to database!")
		}

		db.AutoMigrate(&Agent{})
		db.AutoMigrate(&Command{})
		db.AutoMigrate(&CommandResult{})

	}
	return db
}
