package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"go-web-app-experiment/cmd/app/notebook_db"

	"github.com/gin-contrib/sessions"

	"go-web-app-experiment/cmd/app/user_session"
)

/* Get Requests */

func RenderWebPage(template_page string, page_title string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Get user session data
		session := sessions.Default(c)

		// Get loggin in value
		user := session.Get("is_logged_in")

		c.HTML(http.StatusOK, template_page, gin.H{
			"title": page_title,
			"page_title": page_title,
			"user": user,
		})
	}
	return gin.HandlerFunc(fn)
}

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


