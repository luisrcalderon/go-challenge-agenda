package config

import (
	"os"
)

type Config struct {
	GRPCAddr string
	DBDriver string // "sqlite3" or "postgres"
	DBSource string
}

func Load() Config {
	c := Config{
		GRPCAddr: getenv("AGENDA_GRPC_ADDR", ":50051"),
		DBDriver: getenv("DB_DRIVER", "sqlite3"),
		DBSource: getenv("DB_SOURCE", "agenda.db"),
	}
	return c
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
