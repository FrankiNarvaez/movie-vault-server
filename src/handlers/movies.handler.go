package handlers

import (
	"movie/src/services"
	"movie/src/utils"

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

	utils.HandleResponseOK(c, movies)
}

func GetMovieDetails(c *gin.Context) {
	movie_id := c.Param("movie_id")
	language := c.Query("language")

	movie, err := services.FetchMovieDetails(movie_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, movie)
}

func GetMovieCredits(c *gin.Context) {
	movie_id := c.Param("movie_id")
	language := c.Query("language")

	credits, err := services.FetchMovieCredits(movie_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, credits)
}

func GetMovieImages(c *gin.Context) {
	movie_id := c.Param("movie_id")

	images, err := services.FetchMovieImages(movie_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, images)
}

func GetMovieVideos(c *gin.Context) {
	movie_id := c.Param("movie_id")
	language := c.Query("language")

	videos, err := services.FetchMovieVideos(movie_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, videos)
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

	utils.HandleResponseOK(c, recommendations)
}

func GetMovieExternalIds(c *gin.Context) {
	movie_id := c.Param("movie_id")

	external_ids, err := services.FetchMovieExternalIds(movie_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, external_ids)
}

func GetMovieWatchProviders(c *gin.Context) {
	movie_id := c.Param("movie_id")

	providers, err := services.FetchMovieWatchProviders(movie_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, providers)
}
