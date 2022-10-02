package notebook_db

import (
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"

	"notebook_app/cmd/app/data_types"
)

type User struct {
	gorm.Model
	Email		string
	Username	string
	Password	[]byte
	Note		[]Note
}

type Note struct {
	gorm.Model
	Title       string
	Description string
	UserID		uint
}

/* User functions */

func CreateNewUser(db *gorm.DB, email string, username string, password string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	db.Create(&User{Email: email, Username: username, Password: bytes})
}

func (u *User) Serialize() data_types.JSON {
	return data_types.JSON{
		"id": u.ID,
		"username": u.Username,
	}
}

/* Note functions */

func GetNoteEntries(db *gorm.DB, user_id uint) []Note {
	/* Get all the entries of the notes table */
	var notes []Note // products list

	db.Where("user_id = ?", user_id).Find(&notes)

	//db.Find(&notes)  // find entries of notes table

	return notes
}

func CreateNewNoteEntry(db *gorm.DB, title string, description string, user_id uint) {
	/* Create new note entry */
	db.Create(&Note{Title: title, Description: description, UserID: user_id})
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
