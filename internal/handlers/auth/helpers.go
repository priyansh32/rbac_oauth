package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func generateAuthorizationCode() (string, error) {
	codeBytes := make([]byte, 32)
	_, err := rand.Read(codeBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(codeBytes), nil
}

func verifyhash(value, expectedHash string) bool {
	// Hash the input password using SHA-256
	hashedInputPassword := hasher(value)

	// Compare the hashed input password with the stored hashed password
	return hashedInputPassword == expectedHash
}

func hasher(key string) string {
	// Hash the password using SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hashedkey := hex.EncodeToString(hasher.Sum(nil))
	return hashedkey
}

func generateAccessToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
