package sql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var ErrNoRows = sql.ErrNoRows

func Open(driverName, dataSourceName string) (DB, error) {
	sqlDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &db{db: sqlDB}, nil
}

type db struct {
	db *sql.DB
}

func (d *db) Ping() error {
	return d.db.Ping()
}

func (d *db) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

func (d *db) Query(query string, args ...interface{}) (Rows, error) {
	return d.db.Query(query, args...)
}

func (d *db) QueryRow(query string, args ...interface{}) Row {
	return d.db.QueryRow(query, args...)
}

func (d *db) Close() error {
	return d.db.Close()
}

type DB interface {
	Ping() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row
	Close() error
}

type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
	Err() error
}

type Row interface {
	Scan(dest ...interface{}) error
	Err() error
}
