package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func CelebrityRoutes(api *gin.RouterGroup) {
	celebrities := api.Group("/celebrities")

	celebrities.GET("/popular", handlers.GetPopularCelebrities)

	celebritiesById := celebrities.Group("/:celebrity_id")
	{
		celebritiesById.GET("", handlers.GetCelebrityDetails)
		celebritiesById.GET("/images", handlers.GetCelebrityImages)
		celebritiesById.GET("/movie_credits", handlers.GetCelebrityMovieCredits)
		celebritiesById.GET("/combined_credits", handlers.GetCelebrityCombinedCredits)
		celebritiesById.GET("/external_ids", handlers.GetCelebrityExternalIds)
	}
}
