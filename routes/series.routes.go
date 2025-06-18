package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func SerieRoutes(api *gin.RouterGroup) {
	series := api.Group("/series")

	series.GET("/popular", handlers.GetPopularSeries)

	seriesById := series.Group("/:series_id")
	{
		seriesById.GET("", handlers.GetSeriesDetails)
		seriesById.GET("/credits", handlers.GetSeriesCredits)
		seriesById.GET("/images", handlers.GetSeriesImages)
		seriesById.GET("/videos", handlers.GetSeriesVideos)
		seriesById.GET("/recommendations", handlers.GetSeriesRecommendations)
		seriesById.GET("/external_ids", handlers.GetSeriesExternalIds)
		seriesById.GET("/watch_providers", handlers.GetSeriesWatchProviders)

		seriesBySeasonAndId := seriesById.Group("/season/:season_number")
		{
			seriesBySeasonAndId.GET("", handlers.GetSeasonDetails)
			seriesBySeasonAndId.GET("/credits", handlers.GetSeasonCredits)
			seriesBySeasonAndId.GET("/external_ids", handlers.GetSeasonExternalIds)
			seriesBySeasonAndId.GET("/images", handlers.GetSeasonImages)
			seriesBySeasonAndId.GET("/videos", handlers.GetSeasonVideos)
			seriesBySeasonAndId.GET("/watch_providers", handlers.GetSeasonWatchProviders)

			seriesBySeasonAndIdAndEpisode := seriesBySeasonAndId.Group("/episode/:episode_number")
			{
				seriesBySeasonAndIdAndEpisode.GET("", handlers.GetEpisodeDetails)
				seriesBySeasonAndIdAndEpisode.GET("/credits", handlers.GetEpisodeCredits)
				seriesBySeasonAndIdAndEpisode.GET("/external_ids", handlers.GetEpisodeExternalIds)
				seriesBySeasonAndIdAndEpisode.GET("/images", handlers.GetEpisodeImages)
				seriesBySeasonAndIdAndEpisode.GET("/videos", handlers.GetEpisodeVideos)
			}
		}
	}
}
