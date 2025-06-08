package main

import (
	"log"
	config "movie/internal"
	"movie/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	router := gin.Default()
	routes.RegisterRoutes(router)

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
