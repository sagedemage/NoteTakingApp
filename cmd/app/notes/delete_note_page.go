package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"notebook_app/cmd/app/user_session"
)

func DeleteNoteForm(c *gin.Context) {
	/* Render edit note form */

	// Set title name for the page
	var page_title = "Delete Note"

	// Get user logged_in session data
	user := user_session.GetUserSessionData(c, "is_logged_in")

	// pass the note's title and description to the form page
	c.HTML(http.StatusOK, "delete-note.tmpl", gin.H{
		"page_title":       page_title,
		"user":				user,
	})
}

