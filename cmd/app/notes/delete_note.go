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

		// query note id
		note_id := c.Query("note_id")

		// delete entry
		db.Delete(&notebook_db.Note{}, note_id)

		// Redirect to the dashboard
		c.Redirect(http.StatusFound, "/dashboard")

	}
	return gin.HandlerFunc(fn)
}

