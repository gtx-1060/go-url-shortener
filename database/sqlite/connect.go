package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

var (
	// DSN “file:test.db“ for example
	DSN        = getDsnBase()
	RW_DSN     = DSN + "?_txlock=immediate"
	R_ONLY_DSN = DSN + "?_journal=wal"
)

func getDsnBase() string {
	dsn := os.Getenv("SQLITE_DSN")
	if dsn == "" {
		dsn = "file:TEST.db"
	}
	return dsn
}

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", R_ONLY_DSN)
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
	db, err := sql.Open("sqlite3", RW_DSN)
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
