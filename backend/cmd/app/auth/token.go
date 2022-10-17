package auth

import (
	"io/ioutil"

	"os"

	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"notebook_app/cmd/app/data_types"
)

func generateToken(data data_types.JSON) (string, error) {
	// token is valid for 7 days
	date := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  date.Unix(),
	})

	// get path from root dir
	pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		return "", readErr
	}
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

