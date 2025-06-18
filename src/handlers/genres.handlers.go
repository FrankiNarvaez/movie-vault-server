package handlers

import (
	"movie/src/models"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func GetSerieGenres(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	var genres models.Genres

	_, err := utils.FetchFromTMDB("/genre/tv/list?language="+language, &genres)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, genres)
}

func GetMoviesGenres(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	var genres models.Genres

	_, err := utils.FetchFromTMDB("/genre/movie/list?language="+language, &genres)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, genres)
}
