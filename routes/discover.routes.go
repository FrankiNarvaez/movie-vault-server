package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func DiscoverRoutes(api *gin.RouterGroup) {
	api.GET("/discover/movies", handlers.DiscoverMovies)
	api.GET("/discover/series", handlers.DiscoverSeries)
}
