package auth_middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// verifyToken verifies the JWT token
func verifyToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", token.Header["alg"])
	}

	return token, nil
}
