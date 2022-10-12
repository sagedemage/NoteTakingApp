package notes

import (
	
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/data_types"
)

func ViewNotes(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* View all the database entries as a table */

		type RequestBody struct {
			UserID uint `json:"user_id"`
		}

		var body RequestBody

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



/* Old functions for my purely backend app */
/*
func ViewNotes123(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/ View all the database entries as a table /
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
		/ Handle Delete and Edit Button POST Requests
		on the view notes page /
		// Parse Form Data
		c.Request.ParseForm()

		// initialize query values
		q := url.Values{}

		if form.GetFormValue(c, "delete") != "" {
			/ Delete Note Post request /
			// get entry id for the deleting an entry
			note_id := form.GetFormValue(c, "delete")

			// set note_id query value
			q.Set("note_id", note_id)

			// pass query value to the delete note route
			location := url.URL{Path: "/delete-note", RawQuery: q.Encode()}

			// redirect to th delete note page
			c.Redirect(http.StatusFound, location.RequestURI())

		} else if form.GetFormValue(c, "edit") != "" {
			/ Edit Note Post Request Redirect /
			// get entry id for the editing an entry
			note_id := form.GetFormValue(c, "edit")

			// set note_id query value
			q.Set("note_id", note_id)

			// pass query value to the delete note route
			location := url.URL{Path: "/edit-note", RawQuery: q.Encode()}

			// redirect to edit note
			c.Redirect(http.StatusFound, location.RequestURI())
		}
	}
	return gin.HandlerFunc(fn)
}
*/
