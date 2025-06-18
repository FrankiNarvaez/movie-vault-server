package handlers

import (
	"fmt"
	"movie/src/services"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func GetPopularTvs(c *gin.Context) {
	language := utils.ValidateQueryLanguage(c)
	page := utils.ValidateQueryPage(c)
	movies, err := services.FetchPopularSeries(language, page)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, movies)
}

func GetSeriesDetails(c *gin.Context) {
	series_id := c.Param("series_id")
	language := utils.ValidateQueryLanguage(c)
	append_to_response := c.Query("append_to_response")

	serie, err := services.FetchSeriesDetails(series_id, language, append_to_response)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, serie)
}

func GetSeriesCredits(c *gin.Context) {
	series_id := c.Param("series_id")
	language := utils.ValidateQueryLanguage(c)

	credits, err := services.FetchSeriesCredits(series_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, credits)
}

func GetSeriesImages(c *gin.Context) {
	series_id := c.Param("series_id")

	images, err := services.FetchSeriesImages(series_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, images)
}

func GetSeriesVideos(c *gin.Context) {
	series_id := c.Param("series_id")
	language := utils.ValidateQueryLanguage(c)

	videos, err := services.FetchSeriesVideos(series_id, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, videos)
}

func GetSeriesRecommendations(c *gin.Context) {
	series_id := c.Param("series_id")
	language := utils.ValidateQueryLanguage(c)
	page := utils.ValidateQueryPage(c)

	recommendations, err := services.FetchSeriesRecommendations(series_id, language, page)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, recommendations)
}

func GetSeriesExternalIds(c *gin.Context) {
	series_id := c.Param("series_id")

	external_ids, err := services.FetchSeriesExternalIds(series_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, external_ids)
}

func GetSeriesWatchProviders(c *gin.Context) {
	series_id := c.Param("series_id")

	watch_providers, err := services.FetchSeriesWatchProviders(series_id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, watch_providers)
}

// Season
func GetSeasonDetails(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	language := utils.ValidateQueryLanguage(c)

	season, err := services.FetchSeasonDetails(series_id, season_number, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, season)
}

func GetSeasonCredits(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	language := utils.ValidateQueryLanguage(c)

	credits, err := services.FetchSeasonCredits(series_id, season_number, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, credits)
}

func GetSeasonExternalIds(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")

	external_ids, err := services.FetchSeasonExternalIds(series_id, season_number)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, external_ids)
}

func GetSeasonImages(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")

	images, err := services.FetchSeasonImages(series_id, season_number)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, images)
}

func GetSeasonVideos(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	language := utils.ValidateQueryLanguage(c)

	videos, err := services.FetchSeasonVideos(series_id, season_number, language)
	if err != nil {
		fmt.Println(err)
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, videos)
}

func GetSeasonWatchProviders(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	language := utils.ValidateQueryLanguage(c)

	watch_providers, err := services.FetchSeasonWatchProviders(series_id, season_number, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, watch_providers)
}

func GetEpisodeDetails(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")
	language := utils.ValidateQueryLanguage(c)

	episode, err := services.FetchEpisodeDetails(series_id, season_number, episode_number, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, episode)
}

func GetEpisodeCredits(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")
	language := utils.ValidateQueryLanguage(c)

	credits, err := services.FetchEpisodeCredits(series_id, season_number, episode_number, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, credits)
}

func GetEpisodeExternalIds(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")

	external_ids, err := services.FetchEpisodeExternalIds(series_id, season_number, episode_number)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, external_ids)
}

func GetEpisodeImages(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")

	images, err := services.FetchEpisodeImages(series_id, season_number, episode_number)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, images)
}

func GetEpisodeVideos(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")
	language := utils.ValidateQueryLanguage(c)

	videos, err := services.FetchEpisodeVideos(series_id, season_number, episode_number, language)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleResponseOK(c, videos)
}
