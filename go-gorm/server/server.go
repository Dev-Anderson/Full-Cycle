package server

import (
	"gorm/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server runner and port: %v ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
