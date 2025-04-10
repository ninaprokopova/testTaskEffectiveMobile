package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// DB
	DBhost     string
	DBuser     string
	DBpassword string
	DBname     string
	DBport     string
	DBsslmode  string
	// App
	AppPort string
	// External API
	AgeAPI         string
	GenderAPI      string
	NationalizeAPI string
	// Logs
	LogLevel string
}

const (
	INFO  = "info"
	DEBUG = "debug"
)

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: .env file not found")
	}
	return &Config{
		DBhost:         getEnv("DB_HOST"),
		DBuser:         getEnv("DB_USER"),
		DBpassword:     getEnv("DB_PASSWORD"),
		DBname:         getEnv("DB_NAME"),
		DBport:         getEnv("DB_PORT"),
		DBsslmode:      getEnv("DB_SSLMODE"),
		AppPort:        getEnv("SERVER_PORT"),
		AgeAPI:         getEnv("AGE_API"),
		GenderAPI:      getEnv("GENDER_API"),
		NationalizeAPI: getEnv("NATIONALIZE_API"),
		LogLevel:       getEnv("LOG_LEVEL"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("no value for key %v", key))
	}
	return value
}
