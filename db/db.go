package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	connect()
}

func connect() {
	var err error
	if DB, err = sql.Open("sqlite3", "./db/forum.db"); err != nil {
		log.Fatal(err)
	}
}
