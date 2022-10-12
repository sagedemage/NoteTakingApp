package tests

import (
	"bytes"
	
	"encoding/json"
	
	"io/ioutil"
	
	"net/http"

	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"

	"notebook_app/cmd/app/request_bodies"
)

/* Registration */

func TestRegistrationSuccess(t *testing.T) {
	/* Registration success */

	mockResponse := `{"msg_success":"Registered Successfully","registered":true}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_register := request_bodies.RegisterRequest {
		Email: "test1000@gmail.com",
		Username: "test1000",
		Password: "test1000",
		ConfirmPwd: "test1000",
	}

	jsonValue, _ := json.Marshal(user_register)

	// call register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestRegistrationEmailAlreadyExistsFailure(t *testing.T) {
	/* Registration failure with email that already exists */

	mockResponse := `{"msg_error":"email already taken","registered":false}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_register := request_bodies.RegisterRequest {
		Email: "test1000@gmail.com",
		Username: "test1001",
		Password: "test1001",
		ConfirmPwd: "test1001",
	}

	jsonValue, _ := json.Marshal(user_register)

	// call register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestRegistrationUsernameAlreadyExistsFailure(t *testing.T) {
	/* Registration failure with username that already exists */

	mockResponse := `{"msg_error":"username already taken","registered":false}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_register := request_bodies.RegisterRequest {
		Email: "test1001@gmail.com",
		Username: "test1000",
		Password: "test1001",
		ConfirmPwd: "test1001",
	}

	jsonValue, _ := json.Marshal(user_register)

	// call register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestRegistrationPasswordMustMatchFailure(t *testing.T) {
	/* Registration failure with the passwords not matching */

	mockResponse := `{"msg_error":"passwords do not match","registered":false}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_register := request_bodies.RegisterRequest {
		Email: "test1001@gmail.com",
		Username: "test1001",
		Password: "test1001",
		ConfirmPwd: "test1000",
	}

	jsonValue, _ := json.Marshal(user_register)

	// call register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestRegistrationShortPasswordFailure(t *testing.T) {
	/* Registration failure with short password */

	mockResponse := `{"msg_error":"must be at least 8 characters","registered":false}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_register := request_bodies.RegisterRequest {
		Email: "test100@gmail.com",
		Username: "test100",
		Password: "test100",
		ConfirmPwd: "test100",
	}

	jsonValue, _ := json.Marshal(user_register)

	// call register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}

/* Login */

func TestLoginSuccessWithEmail(t *testing.T) {
	/* Login success with email */

	mockResponse := `{"auth":true,"user_id":1}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_login := request_bodies.LoginRequest {
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
	/* Login success with username */

	mockResponse := `{"auth":true,"user_id":1}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_login := request_bodies.LoginRequest {
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
	/* Login failure with email that does not exists */

	mockResponse := `{"auth":false,"msg_error":"incorrect username or password"}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_login := request_bodies.LoginRequest {
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
	/* Login failure with username that does not exist */

	mockResponse := `{"auth":false,"msg_error":"incorrect username or password"}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_login := request_bodies.LoginRequest {
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
	/* Login failure with incorrect password */

	mockResponse := `{"auth":false,"msg_error":"incorrect username or password"}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	user_login := request_bodies.LoginRequest {
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

// Logout

func TestLogout(t *testing.T) {
	/* Logout success */

	mockResponse := `{"auth":false}`

	var router = Setup() // setup

	// initialize response writer
	writer := httptest.NewRecorder()

	// call login api
	request, _ := http.NewRequest("GET", "/api/logout", nil)
	
	// reponse to an http request
	router.ServeHTTP(writer, request)

	responseData, _ := ioutil.ReadAll(writer.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, writer.Code)
}
