package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	PostgresDSN string
}

func Load() Config {
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "50051"
	}

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN environment variable required")
	}

	return Config{
		Port:        port,
		PostgresDSN: dsn,
	}
}
