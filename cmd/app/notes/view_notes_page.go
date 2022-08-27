package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"go-web-app-experiment/cmd/app/notebook_db"

	"go-web-app-experiment/cmd/app/user_session"
)

/* Get Requests */

/* Notes */

func ViewNotes(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* View all the database entries as a table */
		// Set title name for the page
		var title_name = "Notes"

		// Get User ID Data
		user_id := user_session.GetUserSessionData(c, "user_id").(uint)

		// entries of the notes database
		notes := notebook_db.GetNoteEntries(db, user_id)

		// Get user logged_in session data
		user := user_session.GetUserSessionData(c, "is_logged_in")

		// Pass the list of notes to the web page
		c.HTML(http.StatusOK, "view-notes.tmpl", gin.H{
			"title": title_name,
			"page_title": title_name,
			"note_list":  notes,
			"user": user,
		})
	}
	return gin.HandlerFunc(fn)
}

