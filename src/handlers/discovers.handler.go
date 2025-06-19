package handlers

import (
	"movie/src/services"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func DiscoverMovies(c *gin.Context) {
	includeAdult := utils.ValidateQueryInclude(c, "include_adult")
	includeVideo := utils.ValidateQueryInclude(c, "include_video")
	language := utils.ValidateQueryLanguage(c)
	page := utils.ValidateQueryPage(c)
	sortBy := utils.ValidateQuerySortBy(c)
	withCompanies := c.Query("with_companies")
	withGenres := c.Query("with_genres")

	movies, err := services.FetchDiscoverMovies(includeAdult, includeVideo, language, page, sortBy, withCompanies, withGenres)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, movies)
}

func DiscoverSeries(c *gin.Context) {
	includeAdult := utils.ValidateQueryInclude(c, "include_adult")
	includeNullFirstAirDates := utils.ValidateQueryInclude(c, "include_null_first_air_dates")
	language := utils.ValidateQueryLanguage(c)
	page := utils.ValidateQueryPage(c)
	sortBy := utils.ValidateQuerySortBy(c)
	withCompanies := c.Query("with_companies")
	withGenres := c.Query("with_genres")
	withNetworks := c.Query("with_networks")

	series, err := services.FetchDiscoverSeries(includeAdult, includeNullFirstAirDates, language, page, sortBy, withCompanies, withGenres, withNetworks)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, series)
}
