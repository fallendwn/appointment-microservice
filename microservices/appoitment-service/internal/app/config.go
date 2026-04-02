package app

import (
	"os"
)

type Config struct {
	Port     string
	MongoURI string
}

func NewConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", ":8080"),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
