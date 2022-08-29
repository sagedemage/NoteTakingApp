package notes

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/form"

	"notebook_app/cmd/app/user_session"
)

func AddNewNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Parse Form Data
		c.Request.ParseForm()

		/* Get Title and description from the Post request */
		var title string = form.GetFormValue(c, "title")
		var description string = form.GetFormValue(c, "description")

		// Get User ID Data
		user_id := user_session.GetUserSessionData(c, "user_id").(uint)

		// Create new note entry
		notebook_db.CreateNewNoteEntry(db, title, description, user_id)

		// Redirect to the table view page
		c.Redirect(http.StatusFound, "/view-notes")
	}
	return gin.HandlerFunc(fn)
}

