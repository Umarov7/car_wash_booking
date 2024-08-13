package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	BOOKING_SERVICE_PORT             string
	DB_URI                           string
	DB_NAME                          string
	KAFKA_HOST                       string
	KAFKA_PORT                       string
	KAFKA_TOPIC_BOOKING_CREATED      string
	KAFKA_TOPIC_BOOKING_UPDATED      string
	KAFKA_TOPIC_BOOKING_CANCELLED    string
	KAFKA_TOPIC_PAYMENT_CREATED      string
	KAFKA_TOPIC_REVIEW_CREATED       string
	KAFKA_TOPIC_NOTIFICATION_CREATED string
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

	cfg.KAFKA_HOST = cast.ToString(coalesce("KAFKA_HOST", "localhost"))
	cfg.KAFKA_PORT = cast.ToString(coalesce("KAFKA_PORT", "9092"))

	cfg.KAFKA_TOPIC_BOOKING_CREATED = cast.ToString(coalesce("KAFKA_TOPIC_BOOKING_CREATED", "car-wash:booking_created"))
	cfg.KAFKA_TOPIC_BOOKING_UPDATED = cast.ToString(coalesce("KAFKA_TOPIC_BOOKING_UPDATED", "car-wash:booking_updated"))
	cfg.KAFKA_TOPIC_BOOKING_CANCELLED = cast.ToString(coalesce("KAFKA_TOPIC_BOOKING_CANCELLED", "car-wash:booking_cancelled"))
	cfg.KAFKA_TOPIC_PAYMENT_CREATED = cast.ToString(coalesce("KAFKA_TOPIC_PAYMENT_CREATED", "car-wash:payment_created"))
	cfg.KAFKA_TOPIC_REVIEW_CREATED = cast.ToString(coalesce("KAFKA_TOPIC_REVIEW_CREATED", "car-wash:review_created"))
	cfg.KAFKA_TOPIC_NOTIFICATION_CREATED = cast.ToString(coalesce("KAFKA_TOPIC_NOTIFICATION_CREATED", "car-wash:notification_created"))

	return &cfg
}
