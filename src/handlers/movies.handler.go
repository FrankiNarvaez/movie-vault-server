package handlers

import (
	"movie/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPopularMovies(c *gin.Context) {
	movies, err := services.FetchPopularMovies("en-US")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}

	c.JSON(http.StatusOK, movies)
}
