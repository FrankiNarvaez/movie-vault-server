package routes

import (
	"movie/src/handlers"
	"movie/src/middlewares"

	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(api *gin.RouterGroup) {
	api.GET("/favorites", middlewares.CheckAuth, handlers.GetFavorites)
	api.POST("/favorites", middlewares.CheckAuth, handlers.CreateFavorite)
	api.DELETE("/favorites/:id", middlewares.CheckAuth, handlers.DestroyFavorite)
}
