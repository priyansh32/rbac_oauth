package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return err
	}
	return nil
}

func CloseDatabase() {
	db.Close()
}

func GetResourceOwnerAndType(resourceID string) (string, string, error) {
	var owner string
	var resourceType string
	err := db.QueryRow("SELECT owner, type FROM documents WHERE id=?", resourceID).Scan(&owner, &resourceType)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return "", "", err
	}
	return owner, resourceType, nil
}

func GetDocument(resourceID string) (string, string, error) {
	var title string
	var content string
	err := db.QueryRow("SELECT title, content FROM documents WHERE id=?", resourceID).Scan(&title, &content)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return "", "", err
	}
	return title, content, nil
}

func DeleteDocument(resourceID string) error {
	_, err := db.Exec("DELETE FROM documents WHERE id=?", resourceID)
	if err != nil {
		fmt.Println("Error deleting resource:", err)
		return err
	}
	return nil
}

func UpdateDocument(resourceID, title, content string) error {
	_, err := db.Exec("UPDATE documents SET title=?, content=? WHERE id=?", title, content, resourceID)
	if err != nil {
		fmt.Println("Error updating resource:", err)
		return err
	}
	return nil
}
