package tests

import (
	"net/http"
	"net/url"

	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var r, _ = RunApp()

	// setup http recorder for testing
	write := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("POST", "/register", nil)
	
	v := url.Values{}

	v.Set("email", "test1000@gmail.com")
	v.Add("username", "test1000")
	v.Add("password", "test1000")
	v.Add("confirm", "test1000")

	request.PostForm = v

	r.ServeHTTP(write, request)

	assert.Equal(t, 302, write.Code)

	assert.Equal(t, "/login", write.HeaderMap.Get("Location"))
}

func TestLogin(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var r, _ = RunApp()

	// setup http recorder for testing
	write := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("POST", "/login", nil)
	
	v := url.Values{}

	v.Set("username", "test1000@gmail.com")
	v.Add("password", "test1000")

	request.PostForm = v

	r.ServeHTTP(write, request)

	assert.Equal(t, 302, write.Code)

	assert.Equal(t, "/view-notes", write.HeaderMap.Get("Location"))
}


