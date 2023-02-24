package database

import (
	"gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func GetDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=teste sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Carro{})
}
