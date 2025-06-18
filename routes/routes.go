package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		AuthRoutes(api)
		MoviesRoutes(api)
		CelebrityRoutes(api)
		SerieRoutes(api)
		FavoriteRoutes(api)
		WatchListRoutes(api)
		SearchRoutes(api)
		GenresRoutes(api)
		TrendingRoutes(api)
		DiscoverRoutes(api)
	}
}
