package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatal("Unable to open database: ", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT UNIQUE NOT NULL,
		original_url TEXT NOT NULL

	);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	log.Println("DB Initialized :-)")
}
