package main

import (
	"api-gin-simples/routes"
	"fmt"
)

func main() {
	fmt.Println("Api rodando na porta 8080")
	routes.HandleRequeset()
}
