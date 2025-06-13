package handlers

import (
	"movie/src/services"
	"movie/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPopularMovies(c *gin.Context) {
	language := c.Query("language")
	page := c.Query("page")

	movies, err := services.FetchPopularMovies(language, page)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    movies,
	})
}

func GetTrendingMovies(c *gin.Context) {
	time := c.Param("time_window")
	language := c.Query("language")

	movies, err := services.FetchTrendingMovies(time, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    movies,
	})
}

func GetMovieDetails(c *gin.Context) {
	movie_id := c.Param("movie_id")
	language := c.Query("language")

	movies, err := services.FetchMovieDetails(movie_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetMovieCredits(c *gin.Context) {
	movie_id := c.Param("movie_id")
	language := c.Query("language")

	credtis, err := services.FetchMovieCredits(movie_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, credtis)
}

func GetMovieImages(c *gin.Context) {
	movie_id := c.Param("movie_id")

	images, err := services.FetchMovieImages(movie_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, images)
}

func GetMovieVideos(c *gin.Context) {
	movie_id := c.Param("movie_id")
	language := c.Query("language")

	images, err := services.FetchMovieVideos(movie_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, images)
}

func GetMovieRecommendations(c *gin.Context) {
	movie_id := c.Param("movie_id")
	language := c.Query("language")
	page := c.Query("page")

	recommendations, err := services.FetchMovieRecommendations(movie_id, language, page)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, recommendations)
}

func GetMovieExternalIds(c *gin.Context) {
	movie_id := c.Param("movie_id")

	external_ids, err := services.FetchMovieExternalIds(movie_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, external_ids)
}

func GetMovieWatchProviders(c *gin.Context) {
	movie_id := c.Param("movie_id")

	providers, err := services.FetchMovieWatchProviders(movie_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, providers)
}
