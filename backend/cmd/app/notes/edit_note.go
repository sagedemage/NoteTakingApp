package notes

import (
	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/request_bodies"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
