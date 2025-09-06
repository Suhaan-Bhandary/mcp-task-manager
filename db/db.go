package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DATABASE_FILE_NAME = "task-manager.db"

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DATABASE_FILE_NAME)
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	err := CreateTables(db)
	if err != nil {
		return err
	}

	return nil
}

func CreateTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT NOT NULL,
		created_at INTEGER,
		updated_at INTEGER
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	return nil
}
