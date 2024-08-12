package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	BOOKING_SERVICE_PORT string
	DB_URI               string
	DB_NAME              string
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if exists {
		return value
	}
	return defaultValue
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	cfg := Config{}

	cfg.BOOKING_SERVICE_PORT = cast.ToString(coalesce("BOOKING_SERVICE_PORT", ":50052"))

	cfg.DB_URI = cast.ToString(coalesce("DB_URI", "mongodb://localhost:27017"))
	cfg.DB_NAME = cast.ToString(coalesce("DB_NAME", "test"))

	return &cfg
}
