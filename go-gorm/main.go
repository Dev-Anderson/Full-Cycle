package main

import (
	"gorm/database"
	"gorm/server"
)

func main() {
	database.GetDatabase()
	s := server.NewServer()
	s.Run()
}
