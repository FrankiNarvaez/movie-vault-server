package routes

import (
	"movie/src/handlers"
	"movie/src/middlewares"

	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(api *gin.RouterGroup) {
	favorites := api.Group("favorites")
	favorites.Use(middlewares.CheckAuth)
	{
		favorites.GET("", middlewares.CheckAuth, handlers.GetFavorites)
		favorites.POST("", middlewares.CheckAuth, handlers.CreateFavorite)
		favorites.DELETE("/:id", middlewares.CheckAuth, handlers.DestroyFavorite)
	}
}
