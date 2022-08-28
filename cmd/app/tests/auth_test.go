package tests

import (
	"testing"

	"notebook_app/cmd/app/auth"
)

func TestRegistration(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var _, db = RunApp()

	// get email data
	var email string = "test1000@gmail.com"
		
	// get username form data
	var username string = "test1000"
		
	// get password form data
	var password string = "test1000"

	// get confirm from data
	var confirm string = "test1000"

	err := auth.RegisterNewUser(db, email, username, password, confirm)

	if err != nil {
		t.Fatalf("registration failed")
	}
}

func TestLogin(t *testing.T) {
	// Change directory
	ChangetoRepoRootDirectory()

	// Setup App
	var _, db = RunApp()

	// get username form data
	var username string = "test1000"
		
	// get password form data
	var password string = "test1000"

	_, err := auth.IsUserValid(db, username, password)

	if err != nil {
		t.Fatalf("login failed")
	}
}


