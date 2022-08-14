package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Note struct {
	gorm.Model
	Title string
	Description string
}

func OpenDatabase(database_path string)(*gorm.DB) {
	/* Open the Database */
	db, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}
	
	return db
}

func GetDatabase(db *gorm.DB)([]Note) {
	/* Get all the entries of the Database */
	var notes []Note // products list
	db.Find(&notes) // find entries of products database

	return notes
}
