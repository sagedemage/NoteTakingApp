package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/user_session"

	"notebook_app/cmd/app/form"

	"notebook_app/cmd/app/page_renderer"
)

func EditNoteForm(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Render edit note form */

		// Set title name for the page
		var page_title = "Edit Note"

		// query note id
		note_id := c.Query("note_id")

		// get entry note values
		note := notebook_db.GetNoteEntry(db, note_id)

		// Get user logged_in session data
		user := user_session.GetUserSessionData(c, "is_logged_in")


		if note_id != "" {
			// pass the note's title and description to the form page
			c.HTML(http.StatusOK, "edit-note.tmpl", gin.H{
				"page_title":       page_title,
				"note_title":       note.Title,
				"note_description": note.Description,
				"user":				user,
			})
		} else {
			page_renderer.RenderPageNotFoundWebPage(c)
		}
	}
	return gin.HandlerFunc(fn)
}

func EditNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Edit Note Post Request */
		// Parse Form Data
		c.Request.ParseForm()

		/* Get Title and secription from the Post request */
		var title string = form.GetFormValue(c, "title")
		var description string = form.GetFormValue(c, "description")

		// query note id
		note_id := c.Query("note_id")

		// Update the entry title and description by id
		notebook_db.UpdateNoteEntry(db, note_id, title, description)

		// Redirect to the dashboard
		c.Redirect(http.StatusFound, "/dashboard")

	}
	return gin.HandlerFunc(fn)
}

