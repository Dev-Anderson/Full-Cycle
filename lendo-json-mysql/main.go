package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type testeJson struct {
	json string `json:"json"`
}

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/defense")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var jsonData string
	err = db.QueryRow("SELECT JSON FROM DEFENSE.TESTE WHERE ID = ?", 1).Scan(&jsonData)
	if err != nil {
		panic(err.Error())
	}

	var j testeJson
	err = json.Unmarshal([]byte(jsonData), &j)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Json: %s", j.json)
}
