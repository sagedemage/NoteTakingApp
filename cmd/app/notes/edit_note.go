package notes

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/data_types"
)

func FetchNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Fetch the note data */
		type RequestBody struct {
			NoteID uint `json:"note_id"`
		}

		var body RequestBody

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
			"note":  note,
		})
	}
	return gin.HandlerFunc(fn)
}

// Update Note Post Request

func EditNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Edit Note Post Request */
		type RequestBody struct {
			NoteID uint `json:"note_id"`
			Title string `json:"title"`
			Description string `json:"description"`
		}

		var body RequestBody

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

