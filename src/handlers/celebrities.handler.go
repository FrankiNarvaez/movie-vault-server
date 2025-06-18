package handlers

import (
	"movie/src/services"
	"movie/src/utils"

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

	utils.HandleResponseOK(c, people)
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

func GetPersonDetails(c *gin.Context) {
	person_id := c.Param("person_id")
	language := utils.ValidateQueryLanguage(c)
	person, err := services.FetchPersonDetails(person_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, person)
}

func GetPersonImages(c *gin.Context) {
	person_id := c.Param("person_id")
	images, err := services.FetchPersonImages(person_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, images)
}

func GetPersonMovieCredits(c *gin.Context) {
	person_id := c.Param("person_id")
	language := utils.ValidateQueryLanguage(c)
	credits, err := services.FetchPersonMovieCredits(person_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, credits)
}

func GetPersonCombinedCredits(c *gin.Context) {
	person_id := c.Param("person_id")
	language := utils.ValidateQueryLanguage(c)
	credits, err := services.FetchPersonCombinedCredits(person_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, credits)
}

func GetPersonExternalIds(c *gin.Context) {
	person_id := c.Param("person_id")
	external_ids, err := services.FetchPersonExternalIds(person_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, external_ids)
}
