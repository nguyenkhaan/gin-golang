package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading .env: ", err)
	}
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}