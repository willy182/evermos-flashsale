package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// createDBConnection function for creating database connection
func createDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("postgres", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

// LoadPostgres function for creating database connection Postgres Database
func LoadPostgres() *sql.DB {
	return createDBConnection(os.Getenv("DB"))
}
