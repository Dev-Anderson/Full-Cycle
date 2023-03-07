package routes

import (
	"api-gin-mysql/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		user := main.Group("user")
		{
			user.GET("/", controllers.GetAllUser)
			user.POST("/", controllers.CreateUser)
			user.PATCH("/:id", controllers.UpdateUser)
			user.DELETE("/:id", controllers.UpdateUser)
		}
	}

	return router
}
