package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(database_path string) *gorm.DB {
	/* Setup the Database */

	// Open database
	db, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}

	// Generate notes table structure
	db.AutoMigrate(&Note{})

	return db
}

