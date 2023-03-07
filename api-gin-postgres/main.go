package main

import (
	"api-gin-postgres/database"
	"api-gin-postgres/server"
)

func main() {
	database.Database()

	s := server.NewServer()
	s.Run()
}
