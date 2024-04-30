package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "supersecret"

// User the unique userId and email to generate a JWT token
// And sign the token by using a key
// Same key will be used to check the incoming token later
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
