package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyansh32/rbac_oauth/resource_server/internal/db"
)

// GetResourceHandler handles GET requests to the resource.
func GetUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET user",
	})
}

func GetDocumentHandler(c *gin.Context) {
	title, content, err := db.GetDocument(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "INTERNAL SERVER ERROR",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Document retrieved successfully",
		"title":   title,
		"content": content,
	})
}

func CreateDocumentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST document",
	})
}

func UpdateDocumentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT document",
	})
}

func DeleteDocumentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE document",
	})
}
