package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI     string
	Port      string
	SecretKey string
}

func LoadConfig() Config {
	godotenv.Load(".env")
	return Config{
		DBURI:     getEnv("DB_URI", "postgres://postgres:postgres@localhost:5432/e-commerce"),
		Port:      getEnv("PORT", "3000"),
		SecretKey: getEnv("SECRET_KEY", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
