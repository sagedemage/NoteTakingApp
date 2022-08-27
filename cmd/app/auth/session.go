package auth

import (
	"strconv"

	"math/rand"
)

// Not post
func GenerateSessionToken() string {
	// Warning: do not use in production
	return strconv.FormatInt(rand.Int63(), 16)
}

