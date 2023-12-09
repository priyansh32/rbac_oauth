package auth_middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var actions = map[string]string{
	"GET":    "read",
	"POST":   "write",
	"PUT":    "write",
	"DELETE": "delete",
}

// TokenToContextMiddleware extracts token claims and adds them to the context.
func TokenToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORIZED: " + err.Error()})
			c.Abort()
			return
		}

		token, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORIZED: " + err.Error()})
			c.Abort()
			return
		}

		claims, err := extractUserClaims(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR: " + err.Error()})
			c.Abort()
			return
		}

		// Add token claims to context
		c.Set("Claims", claims)

		// Proceed to the next middleware or handler
		c.Next()
	}
}

// RBACMiddleware enforces RBAC policies based on OPA.
func RBACMiddleware(opaPolicy string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("Claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORIZED: TOKEN CLAIMS MISSING"})
			c.Abort()
			return
		}

		claimsMap, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR: 1"})
			c.Abort()
			return
		}

		userRole, err := extractUserRole(claimsMap)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "PAYLOAD ERROR: 2"})
			c.Abort()
			return
		}

		owner, resourceType, err := getResourceOwnerAndType(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR: 1"})
			c.Abort()
			return
		}

		debugPrint(owner, claimsMap["sub"], resourceType)

		sub, err := extractSubject(claimsMap)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "PAYLOAD ERROR: 1"})
			c.Abort()
			return
		}

		if owner != sub {
			c.JSON(http.StatusForbidden, gin.H{"error": "JWT SUBJECT NOT VALID FOR REQUESTED RESOURCE"})
			c.Abort()
			return
		}

		allowed, err := checkRBAC(opaPolicy, userRole, c.Request.Method, strings.ToLower(resourceType))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ACCESS DENIED"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "FORBIDDEN"})
			c.Abort()
			return
		}

		// Proceed to the next middleware or handler
		c.Next()
	}
}
