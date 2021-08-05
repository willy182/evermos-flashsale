package config

import "gorm.io/gorm"

// Config main
type Config struct {
	DB *gorm.DB
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
