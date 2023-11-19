// cmd/main.go

package main

import (
	"fmt"
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/db"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/handler"
	auth "github.com/priyansh32/rbac_oauth/resource_server/pkg/middleware"
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

	// Apply RBAC middleware to all routes
	r.Use(auth.RBACMiddleware(opaPolicy))

	// Define routes
	r.GET("/api/users/:id", handler.GetUserHandler)
	r.GET("/api/documents/:id", handler.GetDocumentHandler)
	r.POST("/api/documents", handler.CreateDocumentHandler)
	r.PUT("/api/documents/:id", handler.UpdateDocumentHandler)
	r.DELETE("/api/documents/:id", handler.DeleteDocumentHandler)

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
