package handlers

import (
	"movie/src/errors"
	"movie/src/services"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func GetAllTrendings(c *gin.Context) {}

func GetTrendingMovies(c *gin.Context) {
	timeWindow, err := utils.ValidateTimeWindow(c.Param("time_window"))
	if err != nil {
		utils.HandleError(c, errors.NewBadRequestError(err.Error()))
		return
	}
	language := utils.ValidateQueryLanguage(c)

	movies, err := services.FetchTrendingMovies(timeWindow, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, movies)
}

func GetTrendingSeries(c *gin.Context) {
	timeWindow, err := utils.ValidateTimeWindow(c.Param("time_window"))
	if err != nil {
		utils.HandleError(c, errors.NewBadRequestError(err.Error()))
	}
	language := utils.ValidateQueryLanguage(c)
	series, err := services.FetchTrendingSeries(language, timeWindow)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, series)
}

func GetTrendingCelebrities(c *gin.Context) {
	timeWindow, err := utils.ValidateTimeWindow(c.Param("time_window"))
	if err != nil {
		utils.HandleError(c, errors.NewBadRequestError(err.Error()))
	}
	language := utils.ValidateQueryLanguage(c)

	celebrities, err := services.FetchTrendingCelebrities(language, timeWindow)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, celebrities)
}
