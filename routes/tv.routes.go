package routes

import "github.com/gin-gonic/gin"

func TvRoutes(api *gin.RouterGroup) {
	tv := api.Group("/tv")

	tv.GET("/popular")
	tv.GET("/trending")

	tvById := tv.Group("/:tv_id")
	{
		tvById.GET("")
		tvById.GET("/credits")
		tvById.GET("/images")
		tvById.GET("/videos")
		tvById.GET("/recommendations")
		tvById.GET("/external_ids")
		tvById.GET("/watch_providers")

		tvBySeasonAndId := tvById.Group("/season/:season_id")
		{
			tvBySeasonAndId.GET("")
			tvBySeasonAndId.GET("/credits")
			tvBySeasonAndId.GET("/external_ids")
			tvBySeasonAndId.GET("/images")
			tvBySeasonAndId.GET("/videos")
			tvBySeasonAndId.GET("/watch_providers")

			tvBySeasonAndIdAndEpisode := tvBySeasonAndId.Group("/episode/:episode_id")
			{
				tvBySeasonAndIdAndEpisode.GET("")
				tvBySeasonAndIdAndEpisode.GET("/credits")
				tvBySeasonAndIdAndEpisode.GET("/external_ids")
				tvBySeasonAndIdAndEpisode.GET("/images")
				tvBySeasonAndIdAndEpisode.GET("/videos")
			}
		}
	}
}
