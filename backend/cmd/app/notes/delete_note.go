package notes

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/request_bodies"
)

func DeleteNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Edit Note Post Request */
		var body request_bodies.DeleteorFetchNoteRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// delete entry
		db.Delete(&notebook_db.Note{}, body.NoteID)
	}
	return gin.HandlerFunc(fn)
}
