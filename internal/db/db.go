package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB represents the PostgreSQL database connection
type DB struct {
	conn *sql.DB
}

// NewDB creates a new DB instance and establishes a connection to the PostgreSQL database
func NewDB(host, port, user, password, dbname string) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		conn: db,
	}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.conn.Close()
}

// Query executes a SQL query and returns the result
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

// Exec executes a SQL statement and returns the number of rows affected
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}

// Insert executes an INSERT statement and returns the ID of the inserted row
func (db *DB) Insert(query string, args ...interface{}) (int64, error) {
	result, err := db.conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}