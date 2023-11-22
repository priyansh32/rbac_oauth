package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DBClient *sql.DB

func InitDatabase() error {
	var err error
	DBClient, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	return nil
}

func CloseDatabase() {
	DBClient.Close()
}

func GetResourceOwnerAndType(resourceID string) (string, string, error) {
	var owner string
	var resourceType string
	err := DBClient.QueryRow("SELECT owner, type FROM documents WHERE id=?", resourceID).Scan(&owner, &resourceType)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return "", "", err
	}
	return owner, resourceType, nil
}

func GetDocument(resourceID string) (string, string, error) {
	var title string
	var content string
	err := DBClient.QueryRow("SELECT title, content FROM documents WHERE id=?", resourceID).Scan(&title, &content)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return "", "", err
	}
	return title, content, nil
}

func DeleteDocument(resourceID string) error {
	_, err := DBClient.Exec("DELETE FROM documents WHERE id=?", resourceID)
	if err != nil {
		fmt.Println("Error deleting resource:", err)
		return err
	}
	return nil
}

func UpdateDocument(resourceID, title, content string) error {
	_, err := DBClient.Exec("UPDATE documents SET title=?, content=? WHERE id=?", title, content, resourceID)
	if err != nil {
		fmt.Println("Error updating resource:", err)
		return err
	}
	return nil
}

// check if client already exists
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

// insert client
func InsertClient(clientID, secret, role, redirectURI string) error {
	_, err := DBClient.Exec("INSERT INTO clients (id, secret, role, redirect_uri) VALUES (?, ?)", clientID, secret, role, redirectURI)
	if err != nil {
		return err
	}
	return nil
}

func InsertAuthorizationCode(code, client_id string, user_id int, code_challenge string) error {
	_, err := DBClient.Exec("INSERT INTO authorization_codes (code, client_id, user_id, code_challenge) VALUES (?, ?, ?, ?)", code, client_id, user_id, code_challenge)
	if err != nil {
		return err
	}
	return nil

}
