package tests

import (
	"net/http"
	"net/url"

	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Login(router *gin.Engine) *httptest.ResponseRecorder {
	/* Login user for tests */
	// setup reponse writer
	writer := httptest.NewRecorder()

	// request login page
	request, _ := http.NewRequest("POST", "/login", nil)
	
	// initialize query values for the post form  
	v := url.Values{}

	// set post form data
	v.Set("username", "test1000@gmail.com")
	v.Add("password", "test1000")

	// post form request
	request.PostForm = v

	// reponse to an http request
	router.ServeHTTP(writer, request)

	return writer
}

func TestAddNote(t *testing.T) {
	// setup
	var router = Setup()

	// Login User
	writer := Login(router)

	// test home page
	request, _ := http.NewRequest("POST", "/add-new-note", nil)
	
	// initialize query values for the post form  
	v := url.Values{}

	// set post form data
	v.Set("title", "title")
	v.Add("description", "description")

	// post form request
	request.PostForm = v

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /view-notes route
	assert.Equal(t, "/view-notes", writer.HeaderMap.Get("Location"))
}

func TestEditNote(t *testing.T) {
	// setup
	var router = Setup()

	// Login User
	writer := Login(router)

	// test home page
	request, _ := http.NewRequest("POST", "/edit-note?note_id=1", nil)
	
	// initialize query values for the post form  
	v := url.Values{}

	// set post form data
	v.Set("title", "Number One")
	v.Add("description", "You are number one")

	// post form request
	request.PostForm = v

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /view-notes route
	assert.Equal(t, "/view-notes", writer.HeaderMap.Get("Location"))
}

func TestDeleteNote(t *testing.T) {
	// setup
	var router = Setup()

	// Login User
	writer := Login(router)

	// test home page
	request, _ := http.NewRequest("POST", "/delete-note?note_id=1", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /view-notes route
	assert.Equal(t, "/view-notes", writer.HeaderMap.Get("Location"))
}


