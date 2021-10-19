package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDBLiveTime = 5 * time.Minute

// ConnectSQL create database pool for Postgres
func ConnectSQL(dsl string) (*DB, error) {
	d, err := NewDatabase(dsl)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetMaxIdleConns(maxIdleDBConn)
	d.SetConnMaxLifetime(maxDBLiveTime)

	dbConn.SQL = d
	err = pingDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// pingDB tries to ping the database
func pingDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database for the application
func NewDatabase(dsl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
