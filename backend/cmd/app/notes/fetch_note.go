package notes

import (
	"notebook_app/cmd/app/notebook_db"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Fetch the note data in order to
		easily edit note quickly */

		// Get JSON Request Body
		note_id, err := strconv.Atoi(c.DefaultQuery("id", "0"))

		if err != nil {
			println(err)
			return
		}

		// get entry note values
		note := notebook_db.GetNoteEntry(db, uint(note_id))

		// return note data
		c.JSON(200, gin.H{
			"title":       note.Title,
			"description": note.Description,
		})
	}
	return gin.HandlerFunc(fn)
}
