package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(api *gin.RouterGroup) {
	api.POST("/login")
	api.POST("/register")
	api.DELETE("/logout")
}
