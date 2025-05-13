package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port          string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	RedisAddr     string
	RedisPwd      string
	Auth0Domain   string
	Auth0Audience string
}

var AppConfig *Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system env vars")
	}

	AppConfig = &Config{
		Port:          getEnv("PORT", "8080"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "secret"),
		DBName:        getEnv("DB_NAME", "userdb"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPwd:      getEnv("REDIS_PASSWORD", "secret"),
		Auth0Domain:   getEnv("AUTH0_DOMAIN", "auth0"),
		Auth0Audience: getEnv("AUTH0_AUDIENCE", "https://www.googleapis.com/auth/userinfo.email"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
