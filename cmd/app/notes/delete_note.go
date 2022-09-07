package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/page_renderer"
	
	"notebook_app/cmd/app/user_session"
)

func DeleteNoteForm(c *gin.Context) {
	/* Render edit note form */

	// Set title name for the page
	var page_title = "Delete Note"

	// query note id
	note_id := c.Query("note_id")

	// Get user logged_in session data
	user := user_session.GetUserSessionData(c, "is_logged_in")

	if note_id != "" {
		// pass the note's title and description to the form page
		c.HTML(http.StatusOK, "delete-note.tmpl", gin.H{
			"page_title":       page_title,
			"user":				user,
		})
	} else {
		page_renderer.RenderPageNotFoundWebPage(c)
	}
}

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

