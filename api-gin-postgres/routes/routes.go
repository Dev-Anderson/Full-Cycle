package routes

import (
	"api-gin-postgres/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		aluno := main.Group("aluno")
		{
			aluno.GET("/", controllers.GetAllAlunos)
		}
	}

	return router
}
