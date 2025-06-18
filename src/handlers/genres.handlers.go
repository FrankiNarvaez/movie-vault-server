package handlers

import (
	"movie/src/models"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func GetSerieGenres(c *gin.Context) {
	language := c.Query("language")
	var genres models.Genres

	if language == "" {
		language = "en-US"
	}

	_, err := utils.FetchFromTMDB("/genre/tv/list?language="+language, &genres)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, genres)
}

func GetMoviesGenres(c *gin.Context) {
	language := c.Query("language")
	var genres models.Genres

	if language == "" {
		language = "en-US"
	}

	_, err := utils.FetchFromTMDB("/genre/movie/list?language="+language, &genres)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, genres)
}
