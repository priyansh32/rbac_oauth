package auth

import (
	"net/http"

	"github.com/priyansh32/rbac_oauth/resource_server/internal/db"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	Role        string `json:"role"`
	RedirectURI string `json:"redirect_uri"`
}

// Parse username and password from the request body
var loginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterClient(c *gin.Context) {
	var client Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	// check if client already exists
	clientExists, err := db.ClientAlreadyExists(client.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if clientExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Client already exists",
		})
		return
	}

	// insert client into database
	err = db.InsertClient(client.ID, client.Secret, client.Role, client.RedirectURI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Client registered successfully",
	})

}

// it should redirect back to the client with the authorization code
func Authorize(c *gin.Context) {
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")
	role := c.Query("role")
	code_challenge := c.Query("code_challenge")
	// method is always S256 for simplicity

	// Retrieve client from database
	var client Client

	err := db.DBClient.QueryRow("SELECT id, secret, role, redirect_uri FROM clients WHERE id=?", clientID).Scan(&client.ID, &client.Secret, &client.Role, &client.RedirectURI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid client",
		})
		return
	}

	// Check if redirect URI matches
	if client.RedirectURI != redirectURI {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid redirect URI",
		})
		return
	}

	// Check if role matches
	if client.Role != role {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid role",
		})
		return
	}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
		return
	}

	// Retrieve user from the database based on the provided username
	var password_hash string
	var user_id int
	err = db.DBClient.QueryRow("SELECT id, password_hash FROM users WHERE username = ?", loginCredentials.Username).Scan(&user_id, &password_hash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_grant"})
		return
	}

	// Verify the hashed password
	if !verifyhash(loginCredentials.Password, password_hash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_grant"})
		return
	}

	// Generate authorization code
	code, err := generateAuthorizationCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	// Insert authorization code into database
	err = db.InsertAuthorizationCode(code, client.ID, user_id, code_challenge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong!",
		})
		return
	}

	// Redirect back to client with authorization code
	c.Redirect(http.StatusFound, client.RedirectURI+"?code="+code)
}

func Token(c *gin.Context) {
	code := c.PostForm("code")
	codeVerifier := c.PostForm("code_verifier")

	// Retrieve authorization code from database
	var client_id, user_id int
	var role, code_challenge string
	err := db.DBClient.QueryRow("SELECT client_id, role, user_id, code_challenge FROM authorization_codes inner join clients on authorization_codes.client_id = clients.id WHERE code=?", code).Scan(&client_id, &role, &user_id, &code_challenge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid authorization code",
		})
		return
	}
	// remove authorization code from database
	_, err = db.DBClient.Exec("DELETE FROM authorization_codes WHERE code=?", code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	// Verify code challenge
	if !verifyhash(codeVerifier, code_challenge) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid code verifier",
		})
		return
	}

	// Generate access token
	access_token, err := generateAccessToken(user_id, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": access_token,
	})

}
