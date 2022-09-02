package tests

import (
	"net/http"

	"net/url"

	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

/* Registration */
func TestRegistrationSuccess(t *testing.T) {
	var router = Setup() // setup

	// initialize reponse writer
	writer := httptest.NewRecorder() 

	// test home page
	request, _ := http.NewRequest("POST", "/register", nil)
	  
	v := url.Values{} // initialize url values

	// set post form data
	v.Set("email", "test1000@gmail.com")
	v.Add("username", "test1000")
	v.Add("password", "test1000")
	v.Add("confirm", "test1000")

	request.PostForm = v // post form request
	
	// reponse to an http request
	router.ServeHTTP(writer, request) 

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	var route_path = "/login?msg_success=Registered+Successfully"

	// check if the request redirects to the /login route
	assert.Equal(t, route_path, writer.HeaderMap.Get("Location"))
}

func TestRegistrationEmailAlreadyExistsFailure(t *testing.T) {
	/* Test email already exists */
	var router = Setup() // setup

	// initialize reponse writer
	writer := httptest.NewRecorder() 

	// test home page
	request, _ := http.NewRequest("POST", "/register", nil)
	  
	v := url.Values{} // initialize url values
	
	// set post form data
	v.Set("email", "test1000@gmail.com")
	v.Add("username", "test1001")
	v.Add("password", "test1001")
	v.Add("confirm", "test1001")

	request.PostForm = v // post form request
	
	// reponse to an http request
	router.ServeHTTP(writer, request) 

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	var route_path = "/register?msg_error=email+already+taken"

	// check if the request redirects to the /login route
	assert.Equal(t, route_path, writer.HeaderMap.Get("Location"))
}

func TestRegistrationUsernameAlreadyExistsFailure(t *testing.T) {
	/* Test username already exists */
	var router = Setup() // setup

	// initialize reponse writer
	writer := httptest.NewRecorder() 

	// test home page
	request, _ := http.NewRequest("POST", "/register", nil)
	  
	v := url.Values{} // initialize url values
	
	// set post form data
	v.Set("email", "test1001@gmail.com")
	v.Add("username", "test1000")
	v.Add("password", "test1001")
	v.Add("confirm", "test1001")

	request.PostForm = v // post form request
	
	// reponse to an http request
	router.ServeHTTP(writer, request) 

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	var route_path = "/register?msg_error=username+already+taken"

	// check if the request redirects to the /login route
	assert.Equal(t, route_path, writer.HeaderMap.Get("Location"))
}

func TestRegistrationPasswordMustMatchFailure(t *testing.T) {
	/* Test passwords do not match */
	var router = Setup() // setup

	// initialize reponse writer
	writer := httptest.NewRecorder() 

	// test home page
	request, _ := http.NewRequest("POST", "/register", nil)
	  
	v := url.Values{} // initialize url values
	
	// set post form data
	v.Set("email", "test1001@gmail.com")
	v.Add("username", "test1001")
	v.Add("password", "test1001")
	v.Add("confirm", "test1000")

	request.PostForm = v // post form request
	
	// reponse to an http request
	router.ServeHTTP(writer, request) 

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	var route_path = "/register?msg_error=passwords+do+not+match"

	// check if the request redirects to the /login route
	assert.Equal(t, route_path, writer.HeaderMap.Get("Location"))
}

func TestRegistrationShortPasswordFailure(t *testing.T) {
	/* Test passwords do not match */
	var router = Setup() // setup

	// initialize reponse writer
	writer := httptest.NewRecorder() 

	// test home page
	request, _ := http.NewRequest("POST", "/register", nil)
	  
	v := url.Values{} // initialize url values
	
	// set post form data
	v.Set("email", "test100@gmail.com")
	v.Add("username", "test100")
	v.Add("password", "test100")
	v.Add("confirm", "test100")

	request.PostForm = v // post form request
	
	// reponse to an http request
	router.ServeHTTP(writer, request) 

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	var route_path = "/register?msg_error=must+be+at+least+8+characters"

	// check if the request redirects to the /login route
	assert.Equal(t, route_path, writer.HeaderMap.Get("Location"))
}

/* Login */
func TestLoginSuccessWithEmail(t *testing.T) {
	/* Login with email */
	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("POST", "/login", nil)
	
	v := url.Values{} // initialize url values

	// set post form data
	v.Set("username", "test1000@gmail.com")
	v.Add("password", "test1000")

	request.PostForm = v // post form request

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /dashboard route
	assert.Equal(t, "/dashboard", writer.HeaderMap.Get("Location"))
}

func TestLoginSuccessWithUsername(t *testing.T) {
	/* Login with username */
	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("POST", "/login", nil)
	
	v := url.Values{} // initialize url values

	// set post form data
	v.Set("username", "test1000")
	v.Add("password", "test1000")

	request.PostForm = v // post form request

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /dashboard route
	assert.Equal(t, "/dashboard", writer.HeaderMap.Get("Location"))
}

func TestLoginEmailDoesNotExistFailure(t *testing.T) {
	/* Email does not exist failure */
	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("POST", "/login", nil)
	
	v := url.Values{} // initialize url values

	// set post form data
	v.Set("username", "test1001@gmail.com")
	v.Add("password", "test1000")

	request.PostForm = v // post form request

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /view-notes route
	assert.Equal(t, "/login?msg_error=incorrect+username+or+password", writer.HeaderMap.Get("Location"))
}

func TestLoginUsernameDoesNotExistFailure(t *testing.T) {
	/* Username does not exist failure */
	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("POST", "/login", nil)
	
	v := url.Values{} // initialize url values

	// set post form data
	v.Set("username", "test1001")
	v.Add("password", "test1000")

	request.PostForm = v // post form request

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /view-notes route
	assert.Equal(t, "/login?msg_error=incorrect+username+or+password", writer.HeaderMap.Get("Location"))
}

func TestLoginPasswordIncorrectFailure(t *testing.T) {
	/* Password is incorrect failure */
	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	// test home page
	request, _ := http.NewRequest("POST", "/login", nil)
	
	v := url.Values{} // initialize url values

	// set post form data
	v.Set("username", "test1000@gmail.com")
	v.Add("password", "test1001")

	request.PostForm = v // post form request

	// reponse to an http request
	router.ServeHTTP(writer, request)

	// check if redirection is successful
	assert.Equal(t, 302, writer.Code)

	// check if the request redirects to the /view-notes route
	assert.Equal(t, "/login?msg_error=incorrect+username+or+password", writer.HeaderMap.Get("Location"))
}


