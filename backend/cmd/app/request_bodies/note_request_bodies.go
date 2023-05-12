package request_bodies

/* Note Request Bodies */

type AddNoteRequest struct {
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type EditNoteRequest struct {
	NoteID      uint   `json:"note_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DeleteNoteRequest struct {
	NoteID uint `json:"note_id"`
}

type ViewNotesRequest struct {
	UserID uint `json:"user_id"`
}
