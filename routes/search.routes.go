package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func SearchRoutes(rg *gin.RouterGroup) {
	search := rg.Group("/search")
	{
		search.GET("/all", handlers.SearchAll)
		search.GET("/movies", handlers.SearchMovies)
		search.GET("/series", handlers.SearchSeries)
		search.GET("/celebrities", handlers.SearchCelebrities)
		search.GET("/companies", handlers.SearchCompanies)
	}
}
