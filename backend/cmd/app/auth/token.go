package auth

import (
	"errors"
	//"log"
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/joho/godotenv"

	"notebook_app/cmd/app/data_types"
	//"notebook_app/cmd/app/request_bodies"
)

func getSecretKey()(string){
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	secretkey := os.Getenv("JWT_SECRET")

	return secretkey
}

func GenerateToken(user_id uint) (string, error) {
	//secretkey := "abc123"
	secretkey := getSecretKey()

	var SigningKey = []byte(secretkey)

	// HS256 works

	// EdDSA does not work

	//jwt.SigningMethodHS256.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"auth": true,
		"user_id": user_id,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func decode_token_string(tokenString string)(interface{}, interface{}) {
	secretkey := getSecretKey()

	var SigningKey = []byte(secretkey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method:")
		}
		return SigningKey, nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["auth"], claims["user_id"]
	}

	return nil, nil
}

func GetDecodedToken(c *gin.Context) {
	/* Fetch Decode Token */
	//var body request_bodies.LoginRequest

	type TokenRequest struct {
		Token string
	}

	var body TokenRequest

	// Get JSON Request Body
	err := c.BindJSON(&body)

	if err != nil {
		println(err)
		return
	}

	auth, user_id := decode_token_string(body.Token)

	c.JSON(200, data_types.JSON{
		"auth": auth,
		"user_id": user_id,
	})
}

