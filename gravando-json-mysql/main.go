package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TesteJson struct {
	json string `json:"json"`
}

func database() *sql.DB {
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/defense")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	db := database()
	j := TesteJson{json: `"nome": "anderson"`}
	jsonStr, err := json.Marshal(j)
	if err != nil {
		fmt.Println("Falha no Enconding do JSON:", err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO TESTE(JSON) VALUES(?)")
	if err != nil {
		fmt.Println("Falha ao inserir registro: ", err.Error())
	}

	_, err = stmt.Exec(jsonStr)
	if err != nil {
		panic(err.Error())
	}
}
