package handlers

import (
	"movie/src/services"
	"movie/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPopularPeople(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	people, err := services.FetchPopularPeople(language, page)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, people)
}

func validateTimeWindow(c *gin.Context) string {
	time_window := c.Param("time_window")
	if time_window == "" {
		time_window = "day"
	}

	if time_window != "day" && time_window != "week" {
		time_window = "day"
	}

	return time_window
}

func GetTrendingPeople(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	time_window := validateTimeWindow(c)

	people, err := services.FetchTrendingPeople(language, time_window)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, people)
}

func GetPersonDetails(c *gin.Context) {
	person_id := c.Param("person_id")
	language := utils.ValidateQueryLanguage(c)
	person, err := services.FetchPersonDetails(person_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, person)
}

func GetPersonImages(c *gin.Context) {
	person_id := c.Param("person_id")
	images, err := services.FetchPersonImages(person_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, images)
}

func GetPersonMovieCredits(c *gin.Context) {
	person_id := c.Param("person_id")
	language := utils.ValidateQueryLanguage(c)
	credits, err := services.FetchPersonMovieCredits(person_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, credits)
}

func GetPersonCombinedCredits(c *gin.Context) {
	person_id := c.Param("person_id")
	language := utils.ValidateQueryLanguage(c)
	credits, err := services.FetchPersonCombinedCredits(person_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, credits)
}

func GetPersonExternalIds(c *gin.Context) {
	person_id := c.Param("person_id")
	external_ids, err := services.FetchPersonExternalIds(person_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, external_ids)
}
