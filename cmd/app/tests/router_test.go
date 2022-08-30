package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	/* Test routes for that do not require authentication */
	// setup
	var router = Setup()

	// setup reponse writer
	writer := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("GET", "/", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if the request succeeded
	assert.Equal(t, 200, writer.Code)

	// test about page
	request, _ = http.NewRequest("GET", "/about", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)
	assert.Equal(t, 200, writer.Code)

	// test login page
	request, _ = http.NewRequest("GET", "/login", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if the request succeeded
	assert.Equal(t, 200, writer.Code)

	// test register page
	request, _ = http.NewRequest("GET", "/register", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if the request succeeded
	assert.Equal(t, 200, writer.Code)
}

func TestAuthRoutes(t *testing.T) {
	/* Test routes that require authentication */
	// setup
	var router = Setup()

	// setup reponse writer
	writer := httptest.NewRecorder()

	// test view notes page
	request, _ := http.NewRequest("GET", "/view-notes", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if the request got unauthorized reponse
	assert.Equal(t, 401, writer.Code)

	// test about page
	request, _ = http.NewRequest("GET", "/add-new-note", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if the request got unauthorized reponse
	assert.Equal(t, 401, writer.Code)

	// test login page
	request, _ = http.NewRequest("GET", "/edit-note", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if the request got unauthorized reponse
	assert.Equal(t, 401, writer.Code)

	// test logging out
	request, _ = http.NewRequest("GET", "/logout", nil)

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if the request got unauthorized reponse
	assert.Equal(t, 401, writer.Code)
}

func TestPageNotFoundRoutes(t *testing.T) {
	/* Test page not found */
	// setup
	var router = Setup()

	// setup reponse writer
	writer := httptest.NewRecorder()

	// test page not found
	request, _ := http.NewRequest("GET", "/welcome", nil)

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if the request got page not found reponse
	assert.Equal(t, 404, writer.Code)
}
