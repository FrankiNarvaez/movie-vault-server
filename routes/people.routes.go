package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func PeopleRoutes(api *gin.RouterGroup) {
	people := api.Group("/people")

	people.GET("/popular", handlers.GetPopularPeople)
	people.GET("/trending", handlers.GetTrendingPeople)

	peopleById := people.Group("/:person_id")
	{
		peopleById.GET("", handlers.GetPersonDetails)
		peopleById.GET("/images", handlers.GetPersonImages)
		peopleById.GET("/movie_credits", handlers.GetPersonMovieCredits)
		peopleById.GET("/combined_credits", handlers.GetPersonCombinedCredits)
		peopleById.GET("/external_ids", handlers.GetPersonExternalIds)
	}
}
