package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/db"
)

// GetUserHandler handles GET requests for the user resource.
func GetUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET user",
	})
}

// GetDocumentHandler handles GET requests for retrieving a document by ID.
func GetDocumentHandler(c *gin.Context) {
	// Get document details from the database using the provided ID
	title, content, err := db.GetDocumentByID(c.Param("id"))

	if err != nil {
		// If there's an error, return an internal server error response
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
		})
		return
	}

	// Return the document details in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Document retrieved successfully",
		"title":   title,
		"content": content,
	})
}

// Get User Documents Handler handles GET requests for retrieving all documents owned by a user.
func GetUserDocumentsHandler(c *gin.Context) {
	// Get the UserID from the context

	claims, exists := c.Get("Claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORIZED: TOKEN CLAIMS MISSING"})
		c.Abort()
		return
	}

	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR: UNABLE TO CONVERT CLAIMS TO MAP"})
		c.Abort()
		return
	}

	// Extract the UserID from the claims
	userID, exists := claimsMap["sub"].(int)
	print(userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UNAUTHORIZED: USER ID MISSING"})
		c.Abort()
		return
	}

	// Get the documents owned by the user from the database
	documents, err := db.GetDocumentsByOwner(userID)

	if err != nil {
		// If there's an error, return an internal server error response
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
		})
		return
	}

	// Return the documents in the response
	c.JSON(http.StatusOK, gin.H{
		"message":   "Documents retrieved successfully",
		"documents": documents,
	})
}

// CreateDocumentHandler handles POST requests for creating a new document.
func CreateDocumentHandler(c *gin.Context) {
	// For simplicity, let's assume the document details are provided in the request body
	var document db.Document
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

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

	sub, ok := claimsMap["sub"].(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR: 2"})
		c.Abort()
		return
	}

	// Add the new document to the database using the UserID from the context
	err := db.CreateDocument(document.Type, sub, document.Title, document.Content)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
		})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Document created successfully",
	})
}

// UpdateDocumentHandler handles PUT requests for updating an existing document.
func UpdateDocumentHandler(c *gin.Context) {
	// For simplicity, let's assume the document details are provided in the request body
	var updatedDocument db.Document
	if err := c.ShouldBindJSON(&updatedDocument); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	// Update the document in the database using the provided ID
	err := db.UpdateDocument(c.Param("id"), updatedDocument.Title, updatedDocument.Content)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
		})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Document updated successfully",
	})
}

// DeleteDocumentHandler handles DELETE requests for deleting a document by ID.
func DeleteDocumentHandler(c *gin.Context) {
	// Delete the document from the database using the provided ID
	err := db.DeleteDocument(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
		})
		return
	}

	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Document deleted successfully",
	})
}
