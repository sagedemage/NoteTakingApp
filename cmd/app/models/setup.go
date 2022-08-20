package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase(database_path string) *gorm.DB {
	/* Open the Database */
	db, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

