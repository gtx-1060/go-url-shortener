package models

import (
	"database/sql"
	"time"
)

const sqliteSchema = `
CREATE TABLE IF NOT EXISTS m_user (
	Id INTEGER NOT NULL PRIMARY KEY,
	Name VARCHAR(32) UNIQUE,
	Created DATETIME,
	Active BOOLEAN
);

CREATE TABLE IF NOT EXISTS Url (
	Id VARCHAR(64) NOT NULL PRIMARY KEY,
	user_id INTEGER,
	Url STRING,
	Created DATETIME,
	Active BOOLEAN,
	FOREIGN KEY (user_id) REFERENCES m_user(Id)
);
`

type User struct {
	Id      uint64
	Name    string
	Created time.Time
	Active  bool
}

type Url struct {
	Id      string
	UserId  uint64
	Url     string
	Created time.Time
	Active  bool
}

func CreateTables(db *sql.DB) {
	if _, err := db.Exec(sqliteSchema); err != nil {
		panic(err)
	}
}
