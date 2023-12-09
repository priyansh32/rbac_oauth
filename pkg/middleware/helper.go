package auth_middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/db"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/rbac"
)

// verifyToken verifies the JWT token
func verifyToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
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

func extractTokenFromHeader(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("TOKEN MISSING")
	}
	return tokenString[len("Bearer "):], nil
}

func extractUserClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR: 1")
	}
	return claims, nil
}

func extractUserRole(claims jwt.MapClaims) (string, error) {
	userRole, ok := claims["role"].(string)
	if !ok {
		return "", fmt.Errorf("PAYLOAD ERROR: 2")
	}
	return userRole, nil
}

func getResourceOwnerAndType(c *gin.Context) (int, string, error) {
	return db.GetResourceOwnerAndType(c.Param("id"))
}

func debugPrint(owner int, sub interface{}, resourceType string) {
	fmt.Print(owner, " ", sub, " ", resourceType)
}

func extractSubject(claims jwt.MapClaims) (int, error) {
	sub, ok := claims["sub"].(int)
	if !ok {
		return 0, fmt.Errorf("PAYLOAD ERROR: 1")
	}
	return sub, nil
}

func checkRBAC(opaPolicy, userRole, method, resourceType string) (bool, error) {
	return rbac.CheckRBAC(opaPolicy, userRole, actions[method], resourceType)
}
