package lib

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoader() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvPort() string {
	EnvLoader()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3030"
	}
	return PORT
}

func EnvDatabaseUrl() string {
	EnvLoader()
	DatabaseUrl := os.Getenv("DATABASE_URL")
	if DatabaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")

	}
	return DatabaseUrl
}
