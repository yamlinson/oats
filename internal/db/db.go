// Package db provides methods for interacting with the application's data store
package db

import (
	"database/sql"
	"errors"
	"math/rand"
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
	defer db.Close()
	setUpDB(db)
	_, err = db.Exec(`
            INSERT INTO items (
                name,
                list,
                created
            ) VALUES (?, ?, ?)
        `,
		item.Name, item.List, item.Created.Format(time.UnixDate))
	if err != nil {
		return err
	}
	return nil
}

// GetLists returns all lists currently in the database.
func GetLists() (*[]string, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	setUpDB(db)
	rows, err := db.Query(`
            SELECT DISTINCT list FROM items
        `)
	if err != nil {
		return nil, err
	}
	var lists []string
	for rows.Next() {
		var list string
		if err := rows.Scan(&list); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return &lists, nil
}

// GetAllItems returns every item in the database and its associated list
func GetAllItems() (*[]Item, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	setUpDB(db)
	var items []Item
	rows, err := db.Query(`
            SELECT name, list, created FROM items
            ORDER BY
                created ASC,
                list ASC;
        `)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name string
		var list string
		var createdString string
		var created time.Time
		if err := rows.Scan(&name, &list, &createdString); err != nil {
			return nil, err
		}
		created, err = time.Parse(time.UnixDate, createdString)
		if err != nil {
			return nil, err
		}
		item := &Item{
			Name:    name,
			List:    list,
			Created: created,
		}
		items = append(items, *item)
	}
	return &items, nil
}

// GetAllItemsByList returns all items associated with a given list
func GetAllItemsByList(list string) (*[]Item, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	setUpDB(db)
	var items []Item
	rows, err := db.Query(`
                SELECT name, list, created FROM items
                WHERE list = ?
                ORDER BY created ASC
            `, list)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name string
		var list string
		var createdString string
		var created time.Time
		if err := rows.Scan(&name, &list, &createdString); err != nil {
			return nil, err
		}
		created, err = time.Parse(time.UnixDate, createdString)
		if err != nil {
			return nil, err
		}
		item := &Item{
			Name:    name,
			List:    list,
			Created: created,
		}
		items = append(items, *item)
	}
	return &items, nil
}

// GetItem returns an item from a specified or random list
// The item can be the first or last in the queue or chosen randomly
func GetItem(list string, random bool, last bool) (*Item, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	setUpDB(db)
	var item *Item
	// Get first from list
	if len(list) > 0 && !random && !last {
		rows, err := db.Query(`
                SELECT name, list, created FROM items
                WHERE list = ?
                ORDER BY created ASC
                LIMIT 1
            `,
			list)
		if err != nil {
			return nil, err
		}
		_ = rows.Next()
		var name string
		var list string
		var createdString string
		var created time.Time
		if err := rows.Scan(&name, &list, &createdString); err != nil {
			return nil, err
		}
		created, err = time.Parse(time.UnixDate, createdString)
		if err != nil {
			return nil, err
		}
		item = &Item{
			Name:    name,
			List:    list,
			Created: created,
		}
		return item, nil
	}
	// Get last from list
	if len(list) > 0 && !random && last {
		rows, err := db.Query(`
                SELECT name, list, created FROM items
                WHERE list = ?
                ORDER BY created DESC
                LIMIT 1
            `,
			list)
		if err != nil {
			return nil, err
		}
		_ = rows.Next()
		var name string
		var list string
		var createdString string
		var created time.Time
		if err := rows.Scan(&name, &list, &createdString); err != nil {
			return nil, err
		}
		item = &Item{
			Name:    name,
			List:    list,
			Created: created,
		}
		return item, nil
	}
	// Get random from specified list
	if len(list) > 0 && random {
		items, err := GetAllItemsByList(list)
		if err != nil {
			return nil, err
		}
		item = &(*items)[rand.Intn(len(*items))]
		return item, nil
	}
	// Get random from any list
	if len(list) == 0 && random {
		items, err := GetAllItems()
		if err != nil {
			return nil, err
		}
		item = &(*items)[rand.Intn(len(*items))]
		return item, nil
	}
	// Fail out chaotically
	return nil, errors.New("bad options")
}

func openDB() (*sql.DB, error) {
	dbPath := data.DataDir + "oats.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setUpDB(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS items (
            id INTEGER NOT NULL PRIMARY KEY,
            name TEXT,
            list TEXT,
            created TEXT,
            UNIQUE(name, list)
        );
    `)
	if err != nil {
		return err
	}
	return nil
}
