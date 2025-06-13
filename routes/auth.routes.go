package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(api *gin.RouterGroup) {
	api.POST("/login", handlers.Login)
	api.POST("/register", handlers.Register)
}
