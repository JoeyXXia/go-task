package config

import (
	"os"
	"strconv"
)

// Config config
type Config struct {
	DatabaseURI string
	ServerPort  int
	JWTSecret   string
}

// Load load from env
func Load() *Config {
	port, err := strconv.Atoi(getEnv("SEVER_PORT", "8080"))
	if err != nil {
		port = 8080
	}

	return &Config{
		DatabaseURI: getEnv("DATABASE_URL", "postgres://"),
		ServerPort:  port,
		JWTSecret:   getEnv("JWT_SECRET", "secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}