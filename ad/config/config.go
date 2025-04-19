package config

import (
	"os"
	"time"
)

type Config struct {
	MongoURI      string
	MongoDatabase string
	Timeout       time.Duration
}

func LoadConfig() *Config {
	return &Config{
		MongoURI:      getEnv("MONGO_URI", "host"),
		MongoDatabase: getEnv("MONGO_DB", "db"),
		Timeout:       30 * time.Second,
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
