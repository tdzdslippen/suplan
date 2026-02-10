package app

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	HTTPAddr    string
	DatabaseURL string
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func Load() (Config, error) {
	_ = godotenv.Load()
	cfg := Config{
		Env:         getenv("APP_ENV", "dev"),
		HTTPAddr:    getenv("HTTP_ADDR", ":8080"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
	if cfg.DatabaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}

	return cfg, nil
}
