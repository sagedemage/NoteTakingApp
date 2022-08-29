package tests

import (
	"net/http"
	"net/url"

	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Login(r *gin.Engine) *httptest.ResponseRecorder {
	write := httptest.NewRecorder()

	// request login page
	request, _ := http.NewRequest("POST", "/login", nil)
	
	v := url.Values{}

	v.Set("username", "test1000@gmail.com")
	v.Add("password", "test1000")

	request.PostForm = v

	r.ServeHTTP(write, request)

	return write
}

func TestAddNote(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var r, _ = RunApp()

	// Login User
	write := Login(r)

	// test home page
	request, _ := http.NewRequest("POST", "/add-new-note", nil)
	
	v := url.Values{}

	v.Set("title", "title")
	v.Add("description", "description")

	request.PostForm = v

	r.ServeHTTP(write, request)

	assert.Equal(t, 302, write.Code)

	assert.Equal(t, "/view-notes", write.HeaderMap.Get("Location"))
}

func TestEditNote(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var r, _ = RunApp()

	// Login User
	write := Login(r)

	// test home page
	request, _ := http.NewRequest("POST", "/edit-note", nil)
	
	v := url.Values{}

	v.Set("title", "Number One")
	v.Add("description", "You are number one")

	request.PostForm = v

	r.ServeHTTP(write, request)

	assert.Equal(t, 302, write.Code)

	assert.Equal(t, "/view-notes", write.HeaderMap.Get("Location"))
}


