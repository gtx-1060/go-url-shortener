package models

import (
	"database/sql"
	"time"
)

const sqliteSchema = `
CREATE TABLE IF NOT EXISTS m_user (
	id INTEGER NOT NULL PRIMARY KEY,
	name CHAR(32),
	created DATETIME,
	active BOOLEAN
);

CREATE TABLE IF NOT EXISTS url (
	id STRING PRIMARY KEY,
	user_id INTEGER,
	url STRING,
	created DATETIME,
	active BOOLEAN,
	FOREIGN KEY (user_id) REFERENCES m_user(id)
);
`

type User struct {
	id      uint64
	name    string
	created time.Time
	active  bool
}

type Url struct {
	id      string
	userId  uint64
	url     string
	created time.Time
	active  bool
}

func CreateTables(db *sql.DB) {
	if _, err := db.Exec(sqliteSchema); err != nil {
		panic(err)
	}
}
