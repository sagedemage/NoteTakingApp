package tests

import (
	"net/http"
	"net/url"

	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	// setup
	var router = Setup()

	// setup reponse writer
	writer := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("POST", "/register", nil)
	
	// initialize query values for the post form  
	v := url.Values{}

	// set post form data
	v.Set("email", "test1000@gmail.com")
	v.Add("username", "test1000")
	v.Add("password", "test1000")
	v.Add("confirm", "test1000")

	// post form request
	request.PostForm = v
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	var route_path = "/login?msg_success=Registered+Successfully"

	// check if the request redirects to the /login route
	assert.Equal(t, route_path, writer.HeaderMap.Get("Location"))
}

func TestLogin(t *testing.T) {
	// setup
	var router = Setup()

	// setup response writer
	writer := httptest.NewRecorder()

	// test home page
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

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /view-notes route
	assert.Equal(t, "/view-notes", writer.HeaderMap.Get("Location"))
}


