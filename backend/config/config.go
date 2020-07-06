// Package config contains all the initialization configurations
// of the web app.
package config

import "time"

// Config contains ServerConfig and DBConfig.
type Config struct {
	Server   ServerConfig
	DB       DBConfig
	MemStore RedisConfig
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

// GetConfig returns the configurations of the app, including
// ServerConfig and DBConfig.
func GetConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         ":8080",
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		DB: DBConfig{
			Dialect:  "postgres",
			Host:     "db",
			Port:     5432,
			Username: "postgres",
			Password: "password",
			Name:     "yarc",
		},
		MemStore: RedisConfig{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		},
	}
}
