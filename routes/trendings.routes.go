package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func TrendingRoutes(api *gin.RouterGroup) {
	api.GET("/trending/all", handlers.GetAllTrendings)
	api.GET("/trending/movies/:time_window", handlers.GetTrendingMovies)
	api.GET("/trending/series/:time_window", handlers.GetTrendingSeries)
	api.GET("/trending/celebrities/:time_window", handlers.GetTrendingCelebrities)
}
