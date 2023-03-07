package config

import (
	"api-gin-mysql/models"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() models.EnvConfig {
	godotenv.Load()
	return models.EnvConfig{
		User:     os.Getenv("USER"),
		Pass:     os.Getenv("PASS"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Database: os.Getenv("DATABASE"),
	}
}
