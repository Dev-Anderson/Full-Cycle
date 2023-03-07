package database

import (
	"api-gin-mysql/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", config.LoadEnv().User, config.LoadEnv().Pass, config.LoadEnv().Host, config.LoadEnv().Port, config.LoadEnv().Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
