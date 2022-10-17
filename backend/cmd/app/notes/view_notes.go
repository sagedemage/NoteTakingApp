package notes

import (
	
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/data_types"

	"notebook_app/cmd/app/request_bodies"
)

func ViewNotes(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* View all the database entries as a table */
		var body request_bodies.ViewNotesRequest

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		} else {
			// entries of the notes database
			notes := notebook_db.GetNoteEntries(db, body.UserID)

			c.JSON(200, data_types.JSON{
				"notes":  notes,
			})
		} 
	}
	return gin.HandlerFunc(fn)
}

