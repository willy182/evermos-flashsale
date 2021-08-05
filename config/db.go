package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	maxOpenConnection     int           = 100
	maxIddleConnection    int           = 50 // 25% / 50% / 75% / 100% of maxConnection
	maxLifetimeConnection time.Duration = 5  // in minutes
)

// createDBConnection function for creating database connection
func createDBConnection(dsn string) *gorm.DB {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	maxOpenCons := os.Getenv("DB_MAX_OPEN_CONS")
	if maxOpenCons != "" {
		maxOpenConnection, _ = strconv.Atoi(maxOpenCons)
	}

	maxIddleCons := os.Getenv("DB_MAX_IDDLE_CONS")
	if maxIddleCons != "" {
		maxIddleConnection, _ = strconv.Atoi(maxIddleCons)
	}

	maxLifetimeCons := os.Getenv("DB_MAX_LIFETIME_CONS")
	if maxLifetimeCons != "" {
		maxLifetimeConnection, _ = time.ParseDuration(fmt.Sprintf("%sm", maxLifetimeCons))
	}

	// set db config
	db, err := gormDB.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenConnection)
	db.SetMaxIdleConns(maxIddleConnection)
	db.SetConnMaxLifetime(maxLifetimeConnection)

	return gormDB
}

// LoadPostgres function for creating database connection Postgres Database
func LoadPostgres() *gorm.DB {
	return createDBConnection(os.Getenv("DB"))
}
