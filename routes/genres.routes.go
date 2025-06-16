package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func GenresRoutes(rg *gin.RouterGroup) {
	rg.GET("/genres/tv", handlers.GetTvGenres)
	rg.GET("/genres/movies", handlers.GetMoviesGenres)
}
