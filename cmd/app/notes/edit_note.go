package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/user_session"

	"notebook_app/cmd/app/form"
)

func EditNoteForm(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Render edit note form */

		// Set title name for the page
		var page_title = "Edit Note"

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// get entry note values
		note := notebook_db.GetNoteEntry(db, id)

		// Get user logged_in session data
		user := user_session.GetUserSessionData(c, "is_logged_in")

		// pass the note's title and description to the form page
		c.HTML(http.StatusOK, "edit-note.tmpl", gin.H{
			"page_title":       page_title,
			"note_title":       note.Title,
			"note_description": note.Description,
			"user":				user,
		})

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

		// Get cookie of the id value
		id, err := c.Cookie("id")
		if err != nil {
			panic(err)
		}

		// Update the entry title and description by id
		notebook_db.UpdateNoteEntry(db, id, title, description)

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")

	}
	return gin.HandlerFunc(fn)
}

