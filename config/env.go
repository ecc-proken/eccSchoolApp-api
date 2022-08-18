package config

import (
	"log"

	"github.com/joho/godotenv"
)

// .envファイルを読み込む
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
