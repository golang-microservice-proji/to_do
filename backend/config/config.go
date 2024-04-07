package config

import (
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg := &Config{
		Port: port,
	}

	return cfg
}

