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

func GetEntry(db *gorm.DB, id string)(*Note) {
	/* Get the entry by id */
	var note = &Note{}
	db.First(&note, id)
	return note
}

func UpdateEntry(db *gorm.DB, id string, title string, description string) {
	/* Update the entry's title and description by id */
	var note = &Note{}

	// Find the first record that matches the id
	db.First(&note, id) 

	// Update Title and Description text
	db.Model(&note).Updates(Note{Title: title, Description: description})
}

