// Package config contains all the initialization configurations
// of the web app.
package config

import (
	"os"
	"strconv"
	"time"
)

// Config contains ServerConfig and DBConfig.
type Config struct {
	Server ServerConfig
	DB     DBConfig
	Redis  RedisConfig
}

// ServerConfig defines the configurations related to the web server.
type ServerConfig struct {
	Port         string
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// DBConfig defines the configurations related to the DB connection.
type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

// RedisConfig defines the configurations related to Redis.
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

var conf Config

// Initialize the config struct from the environment variables.
func init() {
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	rdbNumber, _ := strconv.Atoi(os.Getenv("RDB_DB_NUMBER"))

	conf = Config{
		Server: ServerConfig{
			Port:         ":" + os.Getenv("PORT"),
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		DB: DBConfig{
			Dialect:  os.Getenv("DB_DIALECT"),
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Redis: RedisConfig{
			Addr:     os.Getenv("RDB_ADDR"),
			Password: os.Getenv("RDB_PASSWORD"),
			DB:       rdbNumber,
		},
	}
}

// GetConfig returns the configurations of the app, including
// ServerConfig and DBConfig.
func GetConfig() *Config {
	return &conf
}
