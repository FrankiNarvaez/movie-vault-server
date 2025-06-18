package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func MoviesRoutes(api *gin.RouterGroup) {
	movies := api.Group("/movies")

	movies.GET("/popular", handlers.GetPopularMovies)

	moviesById := movies.Group("/:movie_id")
	{
		moviesById.GET("", handlers.GetMovieDetails)
		moviesById.GET("/images", handlers.GetMovieImages)
		moviesById.GET("/videos", handlers.GetMovieVideos)
		moviesById.GET("/credits", handlers.GetMovieCredits)
		moviesById.GET("/recommendations", handlers.GetMovieRecommendations)
		moviesById.GET("/external_ids", handlers.GetMovieExternalIds)
		moviesById.GET("/watch_providers", handlers.GetMovieWatchProviders)
	}
}
