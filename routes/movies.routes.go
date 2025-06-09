package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func MoviesRoutes(api *gin.RouterGroup) {
	movies := api.Group("/movies")

	movies.GET("/popular", handlers.GetPopularMovies)
	movies.GET("/trending")

	moviesById := movies.Group("/:movie_id")
	{
		moviesById.GET("")
		moviesById.GET("/images")
		moviesById.GET("/videos")
		moviesById.GET("/recommendations")
		moviesById.GET("/external_ids")
		moviesById.GET("/watch_providers")
	}
}
