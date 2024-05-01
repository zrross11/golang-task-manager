package task

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB(dbPath string) (*sql.DB, error) {
	// Open the database file. If it doesn't exist, it will be created.
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create tasks table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        name TEXT,
                        description TEXT
                    )`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
