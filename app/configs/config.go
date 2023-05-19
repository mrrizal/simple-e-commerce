package configs

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI     string
	Port      string
	SecretKey string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	godotenv.Load(".env")
	once.Do(func() {
		config = &Config{
			DBURI:     getEnv("DB_URI", "postgres://postgres:postgres@localhost:5432/e-commerce"),
			Port:      getEnv("PORT", "3000"),
			SecretKey: getEnv("SECRET_KEY", "secret"),
		}
	})
	return config
}
