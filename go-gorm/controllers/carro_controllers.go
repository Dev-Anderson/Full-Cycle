package controllers

import (
	"gorm/database"
	"gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Carro(c *gin.Context) {
	var carro []models.Carro

	database.DB.Find(&carro)
	c.JSON(http.StatusOK, carro)
}
