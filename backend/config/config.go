// Package config contains all the initialization configurations
// of the web app.
package config

import (
	"net/url"
	"os"
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
	Dialect       string
	ConnectionURL string
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
	redisURL, _ := url.Parse(os.Getenv("REDIS_URL"))
	redisPassword, _ := redisURL.User.Password()

	conf = Config{
		Server: ServerConfig{
			Port:         ":" + os.Getenv("PORT"),
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		DB: DBConfig{
			Dialect:       os.Getenv("DATABASE_DIALECT"),
			ConnectionURL: os.Getenv("DATABASE_URL"),
		},
		Redis: RedisConfig{
			Addr:     redisURL.Host,
			Password: redisPassword,
			DB:       0,
		},
	}
}

// GetConfig returns the configurations of the app, including
// ServerConfig and DBConfig.
func GetConfig() *Config {
	return &conf
}
