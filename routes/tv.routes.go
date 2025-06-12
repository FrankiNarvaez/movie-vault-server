package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func TvRoutes(api *gin.RouterGroup) {
	tv := api.Group("/tv")

	tv.GET("/popular", handlers.GetPopularTvs)
	tv.GET("/trending/:time_window", handlers.GetTrendingTvs)

	tvById := tv.Group("/:series_id")
	{
		tvById.GET("", handlers.GetSeriesDetails)
		tvById.GET("/credits", handlers.GetSeriesCredits)
		tvById.GET("/images", handlers.GetSeriesImages)
		tvById.GET("/videos", handlers.GetSeriesVideos)
		tvById.GET("/recommendations", handlers.GetSeriesRecommendations)
		tvById.GET("/external_ids", handlers.GetSeriesExternalIds)
		tvById.GET("/watch_providers", handlers.GetSeriesWatchProviders)

		tvBySeasonAndId := tvById.Group("/season/:season_number")
		{
			tvBySeasonAndId.GET("", handlers.GetSeasonDetails)
			tvBySeasonAndId.GET("/credits", handlers.GetSeasonCredits)
			tvBySeasonAndId.GET("/external_ids", handlers.GetSeasonExternalIds)
			tvBySeasonAndId.GET("/images", handlers.GetSeasonImages)
			tvBySeasonAndId.GET("/videos", handlers.GetSeasonVideos)
			tvBySeasonAndId.GET("/watch_providers", handlers.GetSeasonWatchProviders)

			tvBySeasonAndIdAndEpisode := tvBySeasonAndId.Group("/episode/:episode_number")
			{
				tvBySeasonAndIdAndEpisode.GET("", handlers.GetEpisodeDetails)
				tvBySeasonAndIdAndEpisode.GET("/credits", handlers.GetEpisodeCredits)
				tvBySeasonAndIdAndEpisode.GET("/external_ids", handlers.GetEpisodeExternalIds)
				tvBySeasonAndIdAndEpisode.GET("/images", handlers.GetEpisodeImages)
				tvBySeasonAndIdAndEpisode.GET("/videos", handlers.GetEpisodeVideos)
			}
		}
	}
}
