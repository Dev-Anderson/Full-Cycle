package routes

import (
	"gorm/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		home := main.Group("home")
		{
			home.GET("/")
		}
		carro := main.Group("carro")
		{
			carro.GET("/", controllers.Carro)
		}
	}

	return router
}
