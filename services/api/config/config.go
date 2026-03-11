package config

import "os"

type Config struct {
	HTTPAddr       string
	AgendaGRPCAddr string
}

func Load() Config {
	return Config{
		HTTPAddr:       getenv("API_HTTP_ADDR", ":8080"),
		AgendaGRPCAddr: getenv("AGENDA_GRPC_ADDR", "localhost:50051"),
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
