package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(api *gin.RouterGroup) {
	api.GET("/favorites", handlers.GetFavorites)
	api.POST("/favorites/:imdb_id", handlers.CreateFavorite)
	api.DELETE("/favorites/:imdb_id", handlers.DestroyFavorite)
}
