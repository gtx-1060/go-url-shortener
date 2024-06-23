package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const dsn = "file:test.db"

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
