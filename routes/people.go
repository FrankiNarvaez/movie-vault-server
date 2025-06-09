package routes

import "github.com/gin-gonic/gin"

func PeopleRoutes(api *gin.RouterGroup) {
	people := api.Group("/people")

	people.GET("/popular")
	people.GET("/trending")

	peopleById := people.Group("/:person_id")
	{
		peopleById.GET("")
		peopleById.GET("/images")
		peopleById.GET("/movie_credits")
		peopleById.GET("/combined_credits")
		peopleById.GET("/external_ids")
	}
}
