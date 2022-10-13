package notes

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"notebook_app/cmd/app/notebook_db"
	
	"notebook_app/cmd/app/request_bodies"
)

func AddNewNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Add new note */
		var body request_bodies.AddNoteRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// Create new note entry
		notebook_db.CreateNewNoteEntry(db, body.Title, body.Description, body.UserID)
	}
	return gin.HandlerFunc(fn)
}

