package handlers

import (
	"movie/src/services"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func GetPopularCelebrities(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	celebrities, err := services.FetchPopularCelebrities(language, page)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, celebrities)
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

func GetCelebrityDetails(c *gin.Context) {
	celebrity_id := c.Param("celebrity_id")
	language := utils.ValidateQueryLanguage(c)
	celebrity, err := services.FetchCelebrityDetails(celebrity_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, celebrity)
}

func GetCelebrityImages(c *gin.Context) {
	celebrity_id := c.Param("celebrity_id")
	images, err := services.FetchCelebrityImages(celebrity_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, images)
}

func GetCelebrityMovieCredits(c *gin.Context) {
	celebrity_id := c.Param("celebrity_id")
	language := utils.ValidateQueryLanguage(c)
	credits, err := services.FetchCelebrityMovieCredits(celebrity_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, credits)
}

func GetCelebrityCombinedCredits(c *gin.Context) {
	celebrity_id := c.Param("celebrity_id")
	language := utils.ValidateQueryLanguage(c)
	credits, err := services.FetchCelebrityCombinedCredits(celebrity_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, credits)
}

func GetCelebrityExternalIds(c *gin.Context) {
	celebrity_id := c.Param("celebrity_id")
	external_ids, err := services.FetchCelebrityExternalIds(celebrity_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, external_ids)
}
