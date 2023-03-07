package config

import (
	"api-gin-postgres/models"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() models.EnvConfig {
	godotenv.Load()
	return models.EnvConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		Pass:     os.Getenv("PASS"),
		Database: os.Getenv("DATABASE"),
	}
}
