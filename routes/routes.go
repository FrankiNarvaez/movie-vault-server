package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		movies := api.Group("/movies")

		movies.GET("", handlers.GetPopularMovies)
	}
}
