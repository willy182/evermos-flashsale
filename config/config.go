package config

import (
	"database/sql"
)

// Config main
type Config struct {
	DB *sql.DB
}

var conf *Config

// Load config
func Load() *Config {
	if conf == nil {
		conf = new(Config)
		conf.DB = LoadPostgres()
	}

	return conf
}
