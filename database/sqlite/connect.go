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
	_DSN        = os.Getenv("SQLITE_DSN")
	_RW_DSN     = _DSN + "?_txlock=immediate"
	_R_ONLY_DSN = _DSN + "?_journal=wal"
)

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", _R_ONLY_DSN)
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
	db, err := sql.Open("sqlite3", _RW_DSN)
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
