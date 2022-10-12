package notes

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"notebook_app/cmd/app/notebook_db"
)

func AddNewNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// title and description
		type RequestBody struct {
			Title string `json:"title"`
			Description string `json:"description"`
			UserID uint `json:"user_id"`
		}

		var body RequestBody

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

