package database

import (
	"api-gin-postgres/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Database() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.LoadEnv().Host, config.LoadEnv().Port, config.LoadEnv().User, config.LoadEnv().Pass, config.LoadEnv().Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Falha ao conectar com o banco de dados ", err)
	}

	return db
}
