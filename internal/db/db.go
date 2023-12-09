package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DBClient is the global database client
var DBClient *sql.DB

// Document represents the structure of a document
type Document struct {
	ID      int
	Type    string
	Owner   int
	Title   string
	Content string
}

// InitDatabase initializes the SQLite3 database connection
func InitDatabase() error {
	var err error
	DBClient, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	return nil
}

// CloseDatabase closes the database connection
func CloseDatabase() {
	DBClient.Close()
}

// GetResourceOwnerAndType retrieves the owner and type of a document based on its ID
func GetResourceOwnerAndType(resourceID string) (int, string, error) {
	var owner int
	var resourceType string
	err := DBClient.QueryRow("SELECT owner, type FROM documents WHERE id=?", resourceID).Scan(&owner, &resourceType)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return -1, "", err
	}
	return owner, resourceType, nil
}

// CreateDocument inserts a new document into the database
func CreateDocument(documentType string, owner int, title, content string) error {
	_, err := DBClient.Exec("INSERT INTO documents (type, owner, title, content) VALUES (?, ?, ?, ?)",
		documentType, owner, title, content)
	if err != nil {
		fmt.Println("Error creating document:", err)
		return err
	}
	return nil
}

// GetDocumentByID retrieves the title and content of a document based on its ID
func GetDocumentByID(resourceID string) (string, string, error) {
	var title string
	var content string
	err := DBClient.QueryRow("SELECT title, content FROM documents WHERE id=?", resourceID).Scan(&title, &content)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return "", "", err
	}
	return title, content, nil
}

// GetDocumentByOwner retrieves the id, title and content of a document based on its owner
func GetDocumentsByOwner(ownerID int) ([]Document, error) {
	var documents []Document
	rows, err := DBClient.Query("SELECT id, title, content FROM documents WHERE owner=?", ownerID)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var document Document
		err := rows.Scan(&document.ID, &document.Title, &document.Content)
		if err != nil {
			fmt.Println("Error querying database:", err)
			return nil, err
		}
		documents = append(documents, document)
	}
	return documents, nil
}

// DeleteDocument deletes a document from the database based on its ID
func DeleteDocument(resourceID string) error {
	_, err := DBClient.Exec("DELETE FROM documents WHERE id=?", resourceID)
	if err != nil {
		fmt.Println("Error deleting resource:", err)
		return err
	}
	return nil
}

// UpdateDocument updates the title and content of a document in the database based on its ID
func UpdateDocument(resourceID, title, content string) error {
	_, err := DBClient.Exec("UPDATE documents SET title=?, content=? WHERE id=?", title, content, resourceID)
	if err != nil {
		fmt.Println("Error updating resource:", err)
		return err
	}
	return nil
}

// ClientAlreadyExists checks if a client with the given ID already exists in the database
func ClientAlreadyExists(clientID string) (bool, error) {
	var id string
	err := DBClient.QueryRow("SELECT id FROM clients WHERE id=?", clientID).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		fmt.Println("Error querying database:", err)
		return false, err
	}
	return true, nil
}

// InsertClient inserts a new client into the database
func InsertClient(clientID, secret, role, redirectURI string) error {
	_, err := DBClient.Exec("INSERT INTO clients (id, secret, role, redirect_uri) VALUES (?, ?, ?, ?)", clientID, secret, role, redirectURI)
	if err != nil {
		return err
	}
	return nil
}

// InsertAuthorizationCode inserts a new authorization code into the database
func InsertAuthorizationCode(code, clientID string, userID int, codeChallenge string) error {
	_, err := DBClient.Exec("INSERT INTO authorization_codes (code, client_id, user_id, code_challenge) VALUES (?, ?, ?, ?)", code, clientID, userID, codeChallenge)
	if err != nil {
		return err
	}
	return nil
}
