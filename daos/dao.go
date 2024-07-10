package daos

import (
	"database/sql"
	"log"
)

const (
	defaultRetryTimeoutMs = 200
	defaultMaxRetries     = 5
)

type Tx interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

type Query struct {
	db Tx
}

type RWQuery struct {
	Query
}

type Dao struct {
	concurrentDb    *sql.DB
	nonConcurrentDb *sql.DB
	Query
	retryTimeoutMs uint
	maxRetries     uint
}

/*
func (q Query) QueryRow(query string, args ...any) *sql.Row {
	return q.db.QueryRow(query, args)
}

func (q Query) Exec(query string, args ...any) (sql.Result, error) {
	return q.db.Exec(query, args)
}

func (q Query) Query(query string, args ...any) (*sql.Rows, error) {
	return q.db.Query(query, args)
}
*/

const schema = `
CREATE TABLE IF NOT EXISTS m_user (
	id INTEGER NOT NULL PRIMARY KEY,
	name VARCHAR(32) UNIQUE,
	created DATETIME,
	active BOOLEAN
);

CREATE TABLE IF NOT EXISTS url (
	id VARCHAR(64) NOT NULL PRIMARY KEY,
	user_id INTEGER,
	url VARCHAR(256),
	created DATETIME,
	expiration DATETIME NOT NULL,
	FOREIGN KEY (user_id) REFERENCES m_user(id)
);
`

func (dao *Dao) CreateTables() {
	if _, err := dao.concurrentDb.Exec(schema); err != nil {
		panic(err)
	}
}

func NewDao(concurrentDb *sql.DB, nonConcurrentDb *sql.DB) *Dao {
	return &Dao{
		concurrentDb:    concurrentDb,
		nonConcurrentDb: nonConcurrentDb,
		Query:           Query{concurrentDb},
		retryTimeoutMs:  defaultRetryTimeoutMs,
		maxRetries:      defaultMaxRetries,
	}
}

func (dao *Dao) Destroy() {
	err := dao.concurrentDb.Close()
	if err != nil {
		log.Printf("close db error: %s", err)
	}
	err = dao.nonConcurrentDb.Close()
	if err != nil {
		log.Printf("close nc db error: %s", err)
	}
}
