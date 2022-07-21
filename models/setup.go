package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func OpenDatabase(database_path string)(*gorm.DB) {
	db, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}
	
	return db
}
