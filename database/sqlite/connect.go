package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

// “file:test.db“ for example
func getDSN() string {
	return os.Getenv("SQLITE_DSN")
}

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", getDSN())
	if err != nil {
		log.Fatalf("cant open sqlite db: %s\n", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Duration(3) * time.Minute)
	err = db.Ping()
	if err != nil {
		log.Fatalf("cant open sqlite db: %s\n", err)
	}
	return db
}

func OpenNonConcurrent() *sql.DB {
	db, err := sql.Open("sqlite3", getDSN())
	if err != nil {
		log.Fatalf("cant open sqlite db non-concurrent: %s\n", err)
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(time.Duration(3) * time.Minute)
	err = db.Ping()
	if err != nil {
		log.Fatalf("cant open sqlite db non-concurrent: %s\n", err)
	}
	return db
}
