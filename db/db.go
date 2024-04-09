package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // This is not a direct import, database/sql use this under the hood
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // driver name here is sqlite to use go-sqlite3 under the hood
	// data source name will be any local file name which will be created automatically instead of storing in a real db
	// All the data will be stored in this file and we can interact with it through sql queries.

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

}
