package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title       string
	Description string
}

func GetNoteEntries(db *gorm.DB) []Note {
	/* Get all the entries of the notes table */
	var notes []Note // products list
	db.Find(&notes)  // find entries of notes table

	return notes
}

func GetNoteEntry(db *gorm.DB, id string) *Note {
	/* Get the entry by id */
	var note = &Note{}
	db.First(&note, id)
	return note
}

func UpdateNoteEntry(db *gorm.DB, id string, title string, description string) {
	/* Update the entry's title and description by id */
	var note = &Note{}

	// Find the first record that matches the id
	db.First(&note, id)

	// Update Title and Description text
	db.Model(&note).Updates(Note{Title: title, Description: description})
}
