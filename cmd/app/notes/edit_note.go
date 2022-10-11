package notes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"

	"notebook_app/cmd/app/user_session"

	"notebook_app/cmd/app/form"

	"notebook_app/cmd/app/data_types"
)

// Fetch Note Data

func FetchNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Fetch the note data */
		type RequestBody struct {
			NoteID uint `json:"note_id"`
		}

		var body RequestBody

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// get entry note values
		note := notebook_db.GetNoteEntry(db, body.NoteID)
	
		// return note data
		c.JSON(200, data_types.JSON{
			"note":  note,
		})
	}
	return gin.HandlerFunc(fn)
}

// Update Note Post Request

func EditNote(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Edit Note Post Request */
		type RequestBody struct {
			NoteID uint `json:"note_id"`
			Title string `json:"title"`
			Description string `json:"description"`
		}

		var body RequestBody

		// Get JSON Request Body
		err := c.BindJSON(&body)

		if err != nil {
			println(err)
			return
		}

		// Update the entry title and description by id
		notebook_db.UpdateNoteEntry(db, body.NoteID, body.Title, body.Description)

		// Redirect to the dashboard
		c.Redirect(http.StatusFound, "/dashboard")

	}
	return gin.HandlerFunc(fn)
}



/* Old functions for my purely backend app */

func EditNote123(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Edit Note Post Request */
		// Parse Form Data
		c.Request.ParseForm()

		/* Get Title and secription from the Post request */
		var title string = form.GetFormValue(c, "title")
		var description string = form.GetFormValue(c, "description")

		// query note id
		note_id_int, err := strconv.Atoi(c.Query("note_id"))

		if err != nil {
			println(err)
		}

		note_id := uint(note_id_int)

		// Update the entry title and description by id
		notebook_db.UpdateNoteEntry(db, note_id, title, description)

		// Redirect to the dashboard
		c.Redirect(http.StatusFound, "/dashboard")

	}
	return gin.HandlerFunc(fn)
}

func EditNoteForm123(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		/* Render edit note form */

		// Set title name for the page
		var page_title = "Edit Note"

		// query note id
		note_id_int, err := strconv.Atoi(c.Query("note_id"))

		note_id := uint(note_id_int)
		
		if err != nil {
			println(err)
		}
		
		// get entry note values
		note := notebook_db.GetNoteEntry(db, note_id)

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

