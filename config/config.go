package config

import (
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}
	return os.Getenv(key)
}
