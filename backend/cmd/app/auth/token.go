package auth

import (
	"fmt"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"notebook_app/cmd/app/request_bodies"
)

func getSecretKey()(string){
	/* Get Secret Key from Envionment Variable */
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	secretkey := os.Getenv("JWT_SECRET")

	return secretkey
}

func GenerateToken(user_id uint) (string, error) {
	// create signing key with secretkey
	var SigningKey = []byte(getSecretKey())

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
	// create signing key with secretkey
	var SigningKey = []byte(getSecretKey())

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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
	var body request_bodies.TokenRequest

	// Get JSON Request Body
	err := c.BindJSON(&body)

	if err != nil {
		println(err)
		return
	}

	auth, user_id := decode_token_string(body.Token)

	c.JSON(200, gin.H{
		"auth": auth,
		"user_id": user_id,
	})
}

