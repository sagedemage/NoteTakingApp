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
	// setup
	var router = Setup()

	writer := httptest.NewRecorder()

	add_note := request_bodies.AddNoteRequest {
		UserID: 1,
		Title: "title",
		Description: "description",
	}

	jsonValue, _ := json.Marshal(add_note)

	// test home page
	request, _ := http.NewRequest("POST", "/api/add-new-note", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 200, writer.Code)
}

func TestFetchNote(t *testing.T) {
	/* Add Note */

	mockResponse := `{"description":"description","title":"title"}`

	// setup
	var router = Setup()

	writer := httptest.NewRecorder()

	add_note := request_bodies.DeleteorFetchNoteRequest {
		NoteID: 1,
	}

	jsonValue, _ := json.Marshal(add_note)

	// test home page
	request, _ := http.NewRequest("POST", "/api/fetch-note", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	// check if redirection is successful
	assert.Equal(t, 200, writer.Code)
}

func TestViewNotes(t *testing.T) {
	/* View Notes */

	// setup
	var router = Setup()

	writer := httptest.NewRecorder()

	add_note := request_bodies.ViewNotesRequest {
		UserID: 1,
	}

	jsonValue, _ := json.Marshal(add_note)

	// test home page
	request, _ := http.NewRequest("POST", "/api/view-notes", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 200, writer.Code)
}

func TestEditNote(t *testing.T) {
	/* Edit Note */
	// setup
	var router = Setup()

	writer := httptest.NewRecorder()

	edit_note := request_bodies.EditNoteRequest {
		NoteID: 1,
		Title: "title2",
		Description: "description2",
	}

	jsonValue, _ := json.Marshal(edit_note)

	// test home page
	request, _ := http.NewRequest("POST", "/api/edit-note", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 200, writer.Code)
}

func TestDeleteNote(t *testing.T) {
	/* Delete Note */
	// setup
	var router = Setup()

	writer := httptest.NewRecorder()

	delete_note := request_bodies.DeleteorFetchNoteRequest {
		NoteID: 1,
	}

	jsonValue, _ := json.Marshal(delete_note)

	// test home page
	request, _ := http.NewRequest("POST", "/api/delete-note", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 200, writer.Code)
}

