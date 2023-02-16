package routes

import (
	"api-gin-simples/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequeset() {
	r := gin.Default()
	r.GET("/home", controllers.Home)

	r.Run()
}
