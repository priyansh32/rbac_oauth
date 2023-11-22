// cmd/main.go

package main

import (
	"fmt"
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/db"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/handlers/auth"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/handlers/resource"
	auth_middleware "github.com/priyansh32/rbac_oauth/resource_server/pkg/middleware"
)

func main() {
	r := gin.Default()

	// initialize database
	err := db.InitDatabase()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.CloseDatabase()

	// Load OPA policy
	opaPolicy, err := loadOPAPolicy("policy/main.rego")
	if err != nil {
		panic(err)
	}

	// Load HTML templates
	r.LoadHTMLGlob("templates/*") // Make sure your templates are in the "templates" directory

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	authRoutes := r.Group("/auth")
	{
		// r.POST("/api/register/client")
		authRoutes.GET("/authorize", auth.AuthorizeGET)
		authRoutes.POST("/authorize", auth.AuthorizePOST)
		authRoutes.POST("/token", auth.Token)
		authRoutes.POST("/revoke")
	}

	// RBAC-protected Routes with RBAC middleware
	rbacProtectedRoutes := r.Group("/api")
	rbacProtectedRoutes.Use(auth_middleware.RBACMiddleware(opaPolicy))
	{
		rbacProtectedRoutes.GET("/users/:id", resource.GetUserHandler)
		rbacProtectedRoutes.GET("/documents/:id", resource.GetDocumentHandler)
		rbacProtectedRoutes.POST("/documents", resource.CreateDocumentHandler)
		rbacProtectedRoutes.PUT("/documents/:id", resource.UpdateDocumentHandler)
		rbacProtectedRoutes.DELETE("/documents/:id", resource.DeleteDocumentHandler)
	}

	// Run the server
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func loadOPAPolicy(policyPath string) (string, error) {
	file, err := os.Open(policyPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	policy, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(policy), nil
}
