package controllers

import (
	"api-gin-postgres/database"
	"api-gin-postgres/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAlunos(c *gin.Context) {
	db := database.Database()
	defer db.Close()

	rows, err := db.Query(`
		select 
			id, nome, rg, cpf
		from alunos`)
	if err != nil {
		fmt.Println("Falha na consulta do Aluno ", err)
	}

	var aluno []models.Alunos
	for rows.Next() {
		var Id int
		var Nome string
		var Rg string
		var Cpf string

		err = rows.Scan(&Id, &Nome, &Rg, &Cpf)
		if err != nil {
			fmt.Println("Falha ao consultar Alunos ", err)
		}

		aluno = append(aluno, models.Alunos{Id: Id, Nome: Nome, Rg: Rg, Cpf: Cpf})
	}

	c.JSON(http.StatusOK, aluno)
}
