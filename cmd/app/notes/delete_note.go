package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"
)

func DeleteNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Edit Note Post Request */
		// Parse Form Data
		c.Request.ParseForm()

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// delete entry
		db.Delete(&notebook_db.Note{}, id)

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")

	}
	return gin.HandlerFunc(fn)
}

