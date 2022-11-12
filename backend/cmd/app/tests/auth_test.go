package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"notebook_app/cmd/app/request_bodies"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Registration */

func TestRegistrationSuccess(t *testing.T) {
	/* Registration success */

	// mock response data
	mockResponse := `{"msg_success":"Registered Successfully","registered":true}`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_register := request_bodies.RegisterRequest {
		Email: "test1000@gmail.com",
		Username: "test1000",
		Password: "test1000",
		ConfirmPwd: "test1000",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_register)

	// request for the register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check if the response data is correct
	assert.Equal(t, mockResponse, string(responseData))

	// check if the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegistrationEmailAlreadyExistsFailure(t *testing.T) {
	/* Registration failure with email that already exists */

	// mock reponse data
	mockResponse := `{"msg_error":"email already taken","registered":false}`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_register := request_bodies.RegisterRequest {
		Email: "test1000@gmail.com",
		Username: "test1001",
		Password: "test1001",
		ConfirmPwd: "test1001",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_register)

	// request for register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check the response data is correct
	assert.Equal(t, mockResponse, string(responseData))

	// check the reponse is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegistrationUsernameAlreadyExistsFailure(t *testing.T) {
	/* Registration failure with username that already exists */

	// mock reponse data
	mockResponse := `{"msg_error":"username already taken","registered":false}`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_register := request_bodies.RegisterRequest {
		Email: "test1001@gmail.com",
		Username: "test1000",
		Password: "test1001",
		ConfirmPwd: "test1001",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_register)

	// request for register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get reponse data
	responseData, _ := io.ReadAll(w.Body)

	// check reponse data is correct
	assert.Equal(t, mockResponse, string(responseData))

	// check the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegistrationPasswordMustMatchFailure(t *testing.T) {
	/* Registration failure with the passwords not matching */

	// mock reponse data
	mockResponse := `{"msg_error":"passwords do not match","registered":false}`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_register := request_bodies.RegisterRequest {
		Email: "test1001@gmail.com",
		Username: "test1001",
		Password: "test1001",
		ConfirmPwd: "test1000",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_register)

	// request fot register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get reponse data
	responseData, _ := io.ReadAll(w.Body)

	// check the resonse data is correct
	assert.Equal(t, mockResponse, string(responseData))

	// check the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegistrationShortPasswordFailure(t *testing.T) {
	/* Registration failure with short password */

	// mock resonse data
	mockResponse := `{"msg_error":"must be at least 8 characters","registered":false}`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_register := request_bodies.RegisterRequest {
		Email: "test100@gmail.com",
		Username: "test100",
		Password: "test100",
		ConfirmPwd: "test100",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_register)

	// request for register api
	request, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check response data is correct
	assert.Equal(t, mockResponse, string(responseData))

	// check response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

/* Login */

func TestLoginSuccessWithEmail(t *testing.T) {
	/* Login success with email */

	// mock response data
	mockData := `"auth":true`

	var r = Setup() // router setup

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_login := request_bodies.LoginRequest {
		Username: "test1000@gmail.com",
		Password: "test1000",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_login)

	// request for login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check if the response data is correct
	assert.Contains(t, string(responseData), mockData)

	// check if the reponse is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginSuccessWithUsername(t *testing.T) {
	/* Login success with username */

	// mock response data
	mockData := `"auth":true`

	var r = Setup() // router setup

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_login := request_bodies.LoginRequest {
		Username: "test1000",
		Password: "test1000",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_login)

	// request for login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check if the response data is correct
	assert.Contains(t, string(responseData), mockData)

	// check if the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginEmailDoesNotExistFailure(t *testing.T) {
	/* Login failure with email that does not exists */

	// mock response data
	mockData := `"auth":false`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_login := request_bodies.LoginRequest {
		Username: "test3000@gmail.com",
		Password: "test3000",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_login)

	// call login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check if the response data is correct
	assert.Contains(t, string(responseData), mockData)

	// check if the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginUsernameDoesNotExistFailure(t *testing.T) {
	/* Login failure with username that does not exist */

	// mock response data
	//mockResponse := `{"auth":false,"msg_error":"incorrect username or password"}`
	mockData := `"auth":false`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_login := request_bodies.LoginRequest {
		Username: "test3000",
		Password: "test3000",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_login)

	// call login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check if the response data is correct
	//assert.Equal(t, mockResponse, string(responseData))


	assert.Contains(t, string(responseData), mockData)

	// check if the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginPasswordIncorrectFailure(t *testing.T) {
	/* Login failure with incorrect password */

	// mock response data
	mockResponse := `{"auth":false,"msg_error":"incorrect username or password"}`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_login := request_bodies.LoginRequest {
		Username: "test1000@gmail.com",
		Password: "test2000",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_login)

	// call login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// reponse to an http request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check if the response data is correct
	assert.Equal(t, mockResponse, string(responseData))

	// check if the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDecodedToken(t *testing.T) {
	/* Login success with username */

	// mock response data
	mockResponse := `{"auth":true,"user_id":1}`

	var r = Setup() // router setup

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request body
	user_login := request_bodies.LoginRequest {
		Username: "test1000",
		Password: "test1000",
	}

	// convert to json
	jsonValue, _ := json.Marshal(user_login)

	// request for login api
	request, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	token := string(responseData[22:159])

	// request body
	token_request :=  request_bodies.TokenRequest {
		Token: token,
	}

	// convert to json
	jsonValue, _ = json.Marshal(token_request)

	request, _ = http.NewRequest("POST", "/api/get-decoded-token", bytes.NewBuffer(jsonValue))

	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ = io.ReadAll(w.Body)

	// check if the response data is correct
	assert.Equal(t, string(responseData), mockResponse)

	// check if the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}

// Logout

func TestLogout(t *testing.T) {
	/* Logout success */

	// mock response data
	mockResponse := `{"auth":false}`

	var r = Setup() // setup router

	// writer for the reponse recorder
	w := httptest.NewRecorder()

	// request for login api
	request, _ := http.NewRequest("GET", "/api/logout", nil)
	
	// serve request
	r.ServeHTTP(w, request)

	// get response data
	responseData, _ := io.ReadAll(w.Body)

	// check if the reponse data is correct
	assert.Equal(t, mockResponse, string(responseData))

	// check if the response is a success
	assert.Equal(t, http.StatusOK, w.Code)
}
