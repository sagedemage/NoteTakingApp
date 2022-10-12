package tests

import (
	"bytes"
	
	"encoding/json"
	
	"io/ioutil"
	
	"net/http"

	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

/* Registration */
/* ---
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
	assert.Equal(t, route_path, writer.Header().Get("Location"))
}
--- */
/*
func TestRegistrationEmailAlreadyExistsFailure(t *testing.T) {
	/ Test email already exists /
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
	assert.Equal(t, route_path, writer.Header().Get("Location"))
}

func TestRegistrationUsernameAlreadyExistsFailure(t *testing.T) {
	/ Test username already exists /
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
	assert.Equal(t, route_path, writer.Header().Get("Location"))
}

func TestRegistrationPasswordMustMatchFailure(t *testing.T) {
	/ Test passwords do not match /
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
	assert.Equal(t, route_path, writer.Header().Get("Location"))
}

func TestRegistrationShortPasswordFailure(t *testing.T) {
	/ Test passwords do not match /
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
	assert.Equal(t, route_path, writer.Header().Get("Location"))
}
*/
/* Login */
func TestLoginSuccessWithEmail(t *testing.T) {
	/* Login with email */

	mockResponse := `{"auth":true,"user_id":1}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	user_login := LoginRequest {
		Username: "test1000@gmail.com",
		Password: "test1000",
	}

	jsonValue, _ := json.Marshal(user_login)

	// call login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestLoginSuccessWithUsername(t *testing.T) {
	/* Login with email */

	mockResponse := `{"auth":true,"user_id":1}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	user_login := LoginRequest {
		Username: "test1000",
		Password: "test1000",
	}

	jsonValue, _ := json.Marshal(user_login)

	// call login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestLoginEmailDoesNotExistFailure(t *testing.T) {
	/* Login with email */

	mockResponse := `{"auth":false,"msg_error":"incorrect username or password"}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	user_login := LoginRequest {
		Username: "test3000@gmail.com",
		Password: "test3000",
	}

	jsonValue, _ := json.Marshal(user_login)

	// call login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestLoginUsernameDoesNotExistFailure(t *testing.T) {
	/* Login with email */

	mockResponse := `{"auth":false,"msg_error":"incorrect username or password"}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	user_login := LoginRequest {
		Username: "test3000",
		Password: "test3000",
	}

	jsonValue, _ := json.Marshal(user_login)

	// call login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestLoginPasswordIncorrectFailure(t *testing.T) {
	/* Login with email */

	mockResponse := `{"auth":false,"msg_error":"incorrect username or password"}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	user_login := LoginRequest {
		Username: "test1000@gmail.com",
		Password: "test2000",
	}

	jsonValue, _ := json.Marshal(user_login)

	// call login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

