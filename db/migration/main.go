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
		CREATE TABLE IF NOT EXISTS STUDENTS (
			NIM INTEGER PRIMARY KEY,
			NAME VARCHAR(255) NOT NULL
		);`)

	if err != nil {
		panic(err)
	}
}
