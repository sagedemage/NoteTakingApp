package tests

import (
	"bytes"

	"net/http"

	"net/http/httptest"

	"encoding/json"

	"io/ioutil"
	
	"testing"

	"github.com/stretchr/testify/assert"

	"notebook_app/cmd/app/request_bodies"
)

func TestAddNote(t *testing.T) {
	/* Add Note */
	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	add_note := request_bodies.AddNoteRequest {
		UserID: 1,
		Title: "title",
		Description: "description",
	}

	// convert to json
	jsonValue, _ := json.Marshal(add_note) 

	// request for add new note api
	request, _ := http.NewRequest("POST", "/api/add-new-note", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// check if the response is a success
	assert.Equal(t, 200, w.Code)
}

func TestFetchNote(t *testing.T) {
	/* Add Note */

	// mock response data
	mockResponse := `{"description":"description","title":"title"}`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	add_note := request_bodies.DeleteorFetchNoteRequest {
		NoteID: 1,
	}

	// convert to json
	jsonValue, _ := json.Marshal(add_note)

	// request for fetch note api
	request, _ := http.NewRequest("POST", "/api/fetch-note", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := ioutil.ReadAll(w.Body)

	// check if the response data is correct
	assert.Equal(t, mockResponse, string(responseData))

	// check if the response is a success
	assert.Equal(t, 200, w.Code)
}

func TestViewNotes(t *testing.T) {
	/* View Notes */

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	add_note := request_bodies.ViewNotesRequest {
		UserID: 1,
	}

	// convert to json
	jsonValue, _ := json.Marshal(add_note)

	// request for view notes api
	request, _ := http.NewRequest("POST", "/api/view-notes", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := ioutil.ReadAll(w.Body)

	// check if the response data is correct
	assert.Contains(t, string(responseData), `"Description":"description"`)
	assert.Contains(t, string(responseData), `"Title":"title"`)

	// check if the reponse is a success
	assert.Equal(t, 200, w.Code)
}

func TestEditNote(t *testing.T) {
	/* Edit Note */
	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	edit_note := request_bodies.EditNoteRequest {
		NoteID: 1,
		Title: "title2",
		Description: "description2",
	}

	// convert to json
	jsonValue, _ := json.Marshal(edit_note)

	// request for edit note api
	request, _ := http.NewRequest("POST", "/api/edit-note", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// check if the reponse is a success
	assert.Equal(t, 200, w.Code)
}

func TestDeleteNote(t *testing.T) {
	/* Delete Note */
	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	delete_note := request_bodies.DeleteorFetchNoteRequest {
		NoteID: 1,
	}

	// convert to json
	jsonValue, _ := json.Marshal(delete_note)

	
	// request for delete note api
	request, _ := http.NewRequest("POST", "/api/delete-note", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// check if the response is a success
	assert.Equal(t, 200, w.Code)
}

