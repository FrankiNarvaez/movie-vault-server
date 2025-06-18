package handlers

import (
	"movie/src/services"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func GetTrendingMovies(c *gin.Context) {
	time := c.Param("time_window")
	language := c.Query("language")

	movies, err := services.FetchTrendingMovies(time, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, movies)
}

func GetTrendingSeries(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	time := c.Param("time_window")
	series, err := services.FetchTrendingSeries(language, time)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, series)
}

func GetTrendingCelebrities(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	time_window := validateTimeWindow(c)

	people, err := services.FetchTrendingPeople(language, time_window)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, people)
}
