package utils

import (
	"errors"
	"fmt"
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

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		_, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("Could not parse token")
		return errors.New("Could not parse token.")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("Invalid token!")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("Invalid token claims.")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}
