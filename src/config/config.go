package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		panic("File env cannot be load")
	}

	return os.Getenv(key)
}
