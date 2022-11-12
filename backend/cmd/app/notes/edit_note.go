package notes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"notebook_app/cmd/app/data_types"
	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/request_bodies"
)

func FetchNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Fetch the note data in order to 
		easily edit note quickly */
		var body request_bodies.DeleteorFetchNoteRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// get entry note values
		note := notebook_db.GetNoteEntry(db, body.NoteID)
	
		// return note data
		c.JSON(200, data_types.JSON{
			"title": note.Title,
			"description": note.Description,
		})
	}
	return gin.HandlerFunc(fn)
}

func EditNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Edit Note Post Request */
		var body request_bodies.EditNoteRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// Update the entry title and description by id
		notebook_db.UpdateNoteEntry(db, body.NoteID, body.Title, body.Description)
	}
	return gin.HandlerFunc(fn)
}

