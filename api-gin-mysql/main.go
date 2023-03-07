package main

import (
	"api-gin-mysql/database"
	"api-gin-mysql/server"
)

func main() {
	database.GetDB()

	s := server.NewServer()
	s.Run()
}
