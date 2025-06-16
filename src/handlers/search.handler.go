package handlers

import (
	"movie/src/models"
	"movie/src/services"
	"movie/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateQuery(c *gin.Context) string {
	query := c.Query("query")
	if query == "" {
		query = " "
	}

	query = strings.ReplaceAll(query, " ", "%20")
	return query
}

func SearchAll(c *gin.Context) {
	query := ValidateQuery(c)
	include_adult := utils.ValidateIncludeAdult(c)
	page := utils.ValidateQueryPage(c)
	language := utils.ValidateQueryLanguage(c)

	var filters models.Filters

	if err := c.ShouldBindBodyWithJSON(&filters); err != nil {
	}

	movies, err := services.FetchSearchAll(language, query, include_adult, page, filters)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, movies)
}

func SearchMovies(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	query := ValidateQuery(c)
	include_adult := utils.ValidateIncludeAdult(c)
	page := utils.ValidateQueryPage(c)

	var filters models.Filters

	if err := c.ShouldBindBodyWithJSON(&filters); err != nil {
	}

	movies, err := services.FetchSearchMovies(language, query, include_adult, page, filters)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, movies)
}

func SearchSeries(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	query := ValidateQuery(c)
	include_adult := utils.ValidateIncludeAdult(c)
	page := utils.ValidateQueryPage(c)

	var filters models.Filters

	if err := c.ShouldBindBodyWithJSON(&filters); err != nil {
	}

	series, err := services.FetchSearchSeries(language, query, include_adult, page, filters)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, series)
}

func SearchCelebrities(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	query := ValidateQuery(c)
	include_adult := utils.ValidateIncludeAdult(c)
	page := utils.ValidateQueryPage(c)

	var filters models.Filters

	if err := c.ShouldBindBodyWithJSON(&filters); err != nil {
	}

	celebrities, err := services.FetchSearchCelebrities(language, query, include_adult, page)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, celebrities)
}

func SearchCompanies(c *gin.Context) {
	query := ValidateQuery(c)
	page := utils.ValidateQueryPage(c)

	companies, err := services.FetchSearchCompanies(query, page)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, companies)
}
