package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/user_session"

	"notebook_app/cmd/app/form"
)

func ViewNotes(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* View all the database entries as a table */
		// Set title name for the page
		var page_title = "Notes"

		// Get User ID Data
		user_id := user_session.GetUserSessionData(c, "user_id").(uint)

		// entries of the notes database
		notes := notebook_db.GetNoteEntries(db, user_id)

		// Get user logged_in session data
		user := user_session.GetUserSessionData(c, "is_logged_in")

		// Pass the list of notes to the web page
		c.HTML(http.StatusOK, "view-notes.tmpl", gin.H{
			"page_title": page_title,
			"note_list":  notes,
			"user": user,
		})
	}
	return gin.HandlerFunc(fn)
}

func DeleteOrEditNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Handle Delete and Edit Button POST Requests
		on the view notes page */
		// Parse Form Data
		c.Request.ParseForm()

		if form.GetFormValue(c, "delete") != "" {
			/* Delete Note Post request */
			// get entry id for the deleting an entry
			id := form.GetFormValue(c, "delete")

			// delete entry
			//db.Delete(&notebook_db.Note{}, id)

			// set id of entry to edit
			c.SetCookie("id", id, 10, "/delete-note", c.Request.URL.Hostname(), false, true)

			// redirect to notes view page
			c.Redirect(http.StatusFound, "/delete-note")

		} else if form.GetFormValue(c, "edit") != "" {
			/* Edit Note Post Request Redirect */
			// get entry id for the editing an entry
			id := form.GetFormValue(c, "edit")

			// set id of entry to edit
			c.SetCookie("id", id, 10, "/edit-note", c.Request.URL.Hostname(), false, true)

			// redirect to edit note
			c.Redirect(http.StatusFound, "/edit-note")
		}
	}
	return gin.HandlerFunc(fn)
}

