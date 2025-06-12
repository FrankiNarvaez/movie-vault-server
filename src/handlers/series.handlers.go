package handlers

import (
	"fmt"
	"movie/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateQueryLanguage(c *gin.Context) string {
	language := c.Query("language")
	if language == "" {
		language = "en-US" // default language
	}
	return language
}

func validateQueryPage(c *gin.Context) string {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	return page
}

func GetPopularTvs(c *gin.Context) {
	language := ValidateQueryLanguage(c)
	page := validateQueryPage(c)
	movies, err := services.FetchPopularSeries(language, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetTrendingTvs(c *gin.Context) {
	language := ValidateQueryLanguage(c)
	time := c.Param("time_window")
	movies, err := services.FetchTrendingSeries(language, time)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeriesDetails(c *gin.Context) {
	series_id := c.Param("series_id")
	language := ValidateQueryLanguage(c)
	append_to_response := c.Query("append_to_response")

	movies, err := services.FetchSeriesDetails(series_id, language, append_to_response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeriesCredits(c *gin.Context) {
	series_id := c.Param("series_id")
	language := ValidateQueryLanguage(c)

	movies, err := services.FetchSeriesCredits(series_id, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeriesImages(c *gin.Context) {
	series_id := c.Param("series_id")

	movies, err := services.FetchSeriesImages(series_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeriesVideos(c *gin.Context) {
	series_id := c.Param("series_id")
	language := ValidateQueryLanguage(c)

	movies, err := services.FetchSeriesVideos(series_id, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeriesRecommendations(c *gin.Context) {
	series_id := c.Param("series_id")
	language := ValidateQueryLanguage(c)
	page := validateQueryPage(c)

	movies, err := services.FetchSeriesRecommendations(series_id, language, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeriesExternalIds(c *gin.Context) {
	series_id := c.Param("series_id")

	movies, err := services.FetchSeriesExternalIds(series_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeriesWatchProviders(c *gin.Context) {
	series_id := c.Param("series_id")

	movies, err := services.FetchSeriesWatchProviders(series_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// Season
func GetSeasonDetails(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	language := ValidateQueryLanguage(c)

	movies, err := services.FetchSeasonDetails(series_id, season_number, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeasonCredits(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	language := ValidateQueryLanguage(c)

	movies, err := services.FetchSeasonCredits(series_id, season_number, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeasonExternalIds(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")

	movies, err := services.FetchSeasonExternalIds(series_id, season_number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeasonImages(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")

	movies, err := services.FetchSeasonImages(series_id, season_number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeasonVideos(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	language := ValidateQueryLanguage(c)

	movies, err := services.FetchSeasonVideos(series_id, season_number, language)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetSeasonWatchProviders(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	language := ValidateQueryLanguage(c)

	movies, err := services.FetchSeasonWatchProviders(series_id, season_number, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetEpisodeDetails(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")
	language := ValidateQueryLanguage(c)

	movies, err := services.FetchEpisodeDetails(series_id, season_number, episode_number, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetEpisodeCredits(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")
	language := ValidateQueryLanguage(c)

	episode, err := services.FetchEpisodeCredits(series_id, season_number, episode_number, language)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, episode)
}

func GetEpisodeExternalIds(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")

	episode, err := services.FetchEpisodeExternalIds(series_id, season_number, episode_number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, episode)
}

func GetEpisodeImages(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")

	episode, err := services.FetchEpisodeImages(series_id, season_number, episode_number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, episode)

}

func GetEpisodeVideos(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number := c.Param("season_number")
	episode_number := c.Param("episode_number")
	language := ValidateQueryLanguage(c)

	episode, err := services.FetchEpisodeVideos(series_id, season_number, episode_number, language)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, episode)
}
