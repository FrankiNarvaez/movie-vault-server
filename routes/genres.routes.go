package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func GenresRoutes(rg *gin.RouterGroup) {
	rg.GET("/genres/series", handlers.GetSerieGenres)
	rg.GET("/genres/movies", handlers.GetMoviesGenres)
}
