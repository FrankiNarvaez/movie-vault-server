package handlers

import (
	"movie/src/errors"
	"movie/src/services"
	"movie/src/types"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func filterOutCelebrities(trending *types.TrendingResults) *types.TrendingResults {
	filteredResults := make([]types.TrendingResult, 0)

	for _, item := range trending.Results {
		if item.MediaType == "movie" || item.MediaType == "tv" {
			filteredResults = append(filteredResults, item)
		}
	}

	return &types.TrendingResults{
		Pages:        trending.Pages,
		Results:      filteredResults,
		TotalPages:   trending.TotalPages,
		TotalResults: int64(len(filteredResults)),
	}
}

func GetAllTrendings(c *gin.Context) {
	timeWindow, err := utils.ValidateTimeWindow(c)
	if err != nil {
		utils.HandleError(c, errors.NewBadRequestError("invalid time window"))
		return
	}
	language := utils.ValidateQueryLanguage(c)
	trending, err := services.FetchAllTrendings(timeWindow, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	mediaOutCelebrities := filterOutCelebrities(&trending)

	utils.HandleResponseOK(c, mediaOutCelebrities)
}

func GetTrendingMovies(c *gin.Context) {
	timeWindow, err := utils.ValidateTimeWindow(c)
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
	timeWindow, err := utils.ValidateTimeWindow(c)
	if err != nil {
		utils.HandleError(c, errors.NewBadRequestError(err.Error()))
	}
	language := utils.ValidateQueryLanguage(c)
	series, err := services.FetchTrendingSeries(timeWindow, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, series)
}

func GetTrendingCelebrities(c *gin.Context) {
	timeWindow, err := utils.ValidateTimeWindow(c)
	if err != nil {
		utils.HandleError(c, errors.NewBadRequestError(err.Error()))
	}
	language := utils.ValidateQueryLanguage(c)

	celebrities, err := services.FetchTrendingCelebrities(timeWindow, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, celebrities)
}
