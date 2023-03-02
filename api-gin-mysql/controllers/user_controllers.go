package controllers

import (
	"api-gin-mysql/config"
	"api-gin-mysql/database"
	"api-gin-mysql/entity"
	"api-gin-mysql/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, entity.GetAllUser())
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	entity.CreateUser(user.Name, user.Email)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	_, err := db.Exec(`UPDATE `+config.LoadEnv().Database+`.users
		SET name = ?, email = ? WHERE id = ?`, user.Name, user.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "User alterado com sucesso!",
	})
}

//NAO FUNCIONOU DESSA FORMA
// func UpdateUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User

// 	entity.UpdateUser(user.Name, user.Email, id)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "User alterado com sucesso!",
// 	})
// }

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	_, err := db.Exec(`DELETE FROM `+config.LoadEnv().Database+`.users WHERE id = ?`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "User alterado com sucesso!",
	})
}
