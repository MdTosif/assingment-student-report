package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	Port string
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	config := &Config{
		Port: getEnv("PORT", "8080"),
	}
	
	return config
}

// getEnv gets an environment variable with a fallback default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}