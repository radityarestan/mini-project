package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/mini-project.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Create table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS USER (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(255) NOT NULL
		);`)

	if err != nil {
		panic(err)
	}
}
