package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "user=postgres dbname=teste password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Falha ao conectar com o banco de dados = ", err)
		return
	}

	defer db.Close()

	//Criando usuario 01
	err = create(db, "Anderson da silva", "anderson@gmail.com")
	if err != nil {
		fmt.Println("Falha ao criar usuario: ", err)
		return
	}

	//Criando usuario 02
	err = create(db, "Anderson da silva", "andersondasilva@gmail.com")
	if err != nil {
		fmt.Println("Falha ao criar usuario: ", err)
		return
	}

	//Consultando usuario 01
	name, email, err := read(db, 2)
	if err != nil {
		fmt.Println("Erro ao consulta usuario: ", err)
		return
	}
	fmt.Println("Usuario: ", name, "Email: ", email)

	//Consulta usuario 02
	n, e, _ := read(db, 3)
	fmt.Println("Nome = ", n, "Email = ", e)

	//Alterando usuario
	err = update(db, 1, "Anderson barbosa", "andersonbarbosa@gmail.com")
	if err != nil {
		fmt.Println("Falha ao alterar usuario: ", err)
		return
	}

	//Deletando usuario
	err = delete(db, 6)
}

func create(db *sql.DB, name string, email string) error {
	_, err := db.Exec("INSERT INTO USERS (NAME, EMAIL) VALUES($1, $2)", name, email)
	return err
}

func read(db *sql.DB, id int) (string, string, error) {
	var name string
	var email string
	err := db.QueryRow("SELECT NAME, EMAIL FROM USERS WHERE ID = $1", id).Scan(&name, &email)
	if err != nil {
		return "", "", err
	}
	return name, email, nil
}

func update(db *sql.DB, id int, name string, email string) error {
	_, err := db.Exec("UPDATE USERS SET NAME = $1, EMAIL = $2 WHERE ID = $3", name, email, id)
	return err
}

func delete(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM USERS WHERE ID = $1", id)
	return err
}
