package auth_middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/db"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/rbac"
)

var actions = map[string]string{
	"GET":    "read",
	"POST":   "write",
	"PUT":    "write",
	"DELETE": "delete",
}

// RBACMiddleware enforces RBAC policies based on OPA.
func RBACMiddleware(opaPolicy string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// if header is not present
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORISED: TOKEN MISSING"})
			c.Abort()
			return
		}

		// Assuming the token is in the format "Bearer <token>"
		tokenString = tokenString[len("Bearer "):]

		// Verify token (replace "your-secret-key" with your actual secret key)
		token, err := verifyToken(tokenString, []byte(os.Getenv("ACCESS_TOKEN_SECRET")))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORISED: INVALID TOKEN"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			c.Abort()
			return
		}

		userRole, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "PAYLOAD ERROR"})
			c.Abort()
			return
		}

		owner, resource_type, err := db.GetResourceOwnerAndType(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			c.Abort()
			return
		}

		fmt.Print(owner, " ", claims["sub"], " ", resource_type)

		sub, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "PAYLOAD ERROR"})
			c.Abort()
			return
		}

		if owner != sub {
			c.JSON(http.StatusForbidden, gin.H{"error": "JWT SUBJET NOT VALID FOR REQUESTED RESORCE"})
			c.Abort()
			return
		}

		// Check RBAC using OPA
		allowed, err := rbac.CheckRBAC(opaPolicy, userRole, actions[c.Request.Method], strings.ToLower(resource_type))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ACCESS DENIED"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
