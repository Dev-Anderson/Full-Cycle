package entity

import (
	"api-gin-mysql/config"
	"api-gin-mysql/database"
	"api-gin-mysql/models"
	"fmt"
)

func GetAllUser() []models.UserID {
	db := database.GetDB()
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			id, name, email 
		FROM ` + config.LoadEnv().Database + `.users`)
	if err != nil {
		fmt.Println("Falha ao buscar o users", err)
	}

	var user []models.UserID
	for rows.Next() {
		var id int
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println(err.Error())
		}

		user = append(user, models.UserID{Id: id, Name: name, Email: email})
	}

	return user
}

func CreateUser(name, email string) {
	db := database.GetDB()
	_, err := db.Exec(`
		INSERT INTO `+config.LoadEnv().Database+`.users (name, email) VALUES(?, ?)`, name, email)
	if err != nil {
		fmt.Println("Falha ao inserir Users", err)
	}
}

func UpdateUser(name, email, id string) {
	db := database.GetDB()
	_, err := db.Exec(`
		UPDATE `+config.LoadEnv().Database+`.users
		SET name = ?, email = ?
		WHERE id = ?`, name, email, id)
	if err != nil {
		fmt.Println("Falha ao alterar o Users", err)
	}
}
