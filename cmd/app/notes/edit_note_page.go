package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"go-web-app-experiment/cmd/app/notebook_db"

	"go-web-app-experiment/cmd/app/user_session"
)

func EditNoteForm(orm_db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Render edit note form */

		// Set title name for the page
		var title_name = "Edit Note"

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// get entry note values
		note := notebook_db.GetNoteEntry(orm_db, id)

		// Get user logged_in session data
		user := user_session.GetUserSessionData(c, "is_logged_in")

		// pass the note's title and description to the form page
		c.HTML(http.StatusOK, "edit-note.tmpl", gin.H{
			"title":			title_name,
			"page_title":       title_name,
			"note_title":       note.Title,
			"note_description": note.Description,
			"user":				user,
		})

	}
	return gin.HandlerFunc(fn)
}

