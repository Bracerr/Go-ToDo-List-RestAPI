package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func GetDbParams() DbInitModel {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Ошибка при получении текущей директории: %v", err)
	}

	envFilePath := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(rootDir))), ".env")

	envErr := godotenv.Load(envFilePath)
	if envErr != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", envErr)
	}

	dbInitModel := DbInitModel{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
	}
	return dbInitModel
}
