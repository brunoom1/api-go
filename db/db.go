package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	localDb, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	db = localDb

	if err != nil {
		panic("Failed to connect with database")
	}
}

func GetCurrentConnection() *gorm.DB {
	return db
}
