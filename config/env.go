package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// .envファイルを読み込む
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// .envファイルから取得した値を返すか、デフォルト値を返す
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
