package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var r, _ = RunApp()
	
	// setup http recorder for testing
	write := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 200, write.Code)

	// test about page
	request, _ = http.NewRequest("GET", "/about", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 200, write.Code)

	// test login page
	request, _ = http.NewRequest("GET", "/login", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 200, write.Code)

	// test register page
	request, _ = http.NewRequest("GET", "/register", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 200, write.Code)
}

func TestAuthRoutes(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var r, _ = RunApp()

	// setup http recorder for testing
	write := httptest.NewRecorder()

	// test view notes page
	request, _ := http.NewRequest("GET", "/view-notes", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 401, write.Code)

	// test about page
	request, _ = http.NewRequest("GET", "/add-new-note", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 401, write.Code)

	// test login page
	request, _ = http.NewRequest("GET", "/edit-note", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 401, write.Code)

	// test logging out
	request, _ = http.NewRequest("GET", "/logout", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 401, write.Code)
}

func TestPageNotFoundRoutes(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var r, _ = RunApp()

	// setup http recorder for testing
	write := httptest.NewRecorder()

	// test page not found
	request, _ := http.NewRequest("GET", "/welcome", nil)
	r.ServeHTTP(write, request)
	assert.Equal(t, 404, write.Code)
}
