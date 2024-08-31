// Package db provides methods for interacting with the application's data store
package db

import (
	"database/sql"
	"time"

	"github.com/yamlinson/oats/internal/data"

	// go-sqlite3 is a driver used indirectly via the sql package.
	_ "github.com/mattn/go-sqlite3"
)

// Item describes a to-do item.
type Item struct {
	Name    string
	List    string
	Created time.Time
}

// AddItem adds a given item to the database.
func AddItem(item Item) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer closeDB(db)
	setUpDB(db)
	statement, err := db.Prepare(`
            INSERT INTO items (
                name,
                list,
                created
            ) VALUES (?, ?, ?)
        `)
	if err != nil {
		return err
	}
	statement.Exec(item.Name, item.List, item.Created.String())
	return nil
}

func openDB() (*sql.DB, error) {
	dbPath := data.DataDir + "oats.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func closeDB(db *sql.DB) error {
	return db.Close()
}

func setUpDB(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS items (
            id INTEGER NOT NULL PRIMARY KEY,
            name TEXT UNIQUE,
            list TEXT,
            created TEXT
        );
    `)
	if err != nil {
		return err
	}
	return nil
}
