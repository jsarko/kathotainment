package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBCon *gorm.DB

func InitDB() {
	var err error

	DBCon, err = gorm.Open(sqlite.Open("kathoe.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database.")
	}
}
