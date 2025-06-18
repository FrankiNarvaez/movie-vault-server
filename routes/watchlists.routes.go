package routes

import (
	"movie/src/handlers"
	"movie/src/middlewares"

	"github.com/gin-gonic/gin"
)

func WatchListRoutes(api *gin.RouterGroup) {
	watchlists := api.Group("watchlists")
	{
		watchlists.Use(middlewares.CheckAuth)
		watchlists.GET("", handlers.GetWatchlists)
		watchlists.POST("", handlers.CreateWatchlist)
		watchlists.PATCH("/:watchlist_id", handlers.UpdateWatchlist)
		watchlists.DELETE("/:watchlist_id", handlers.DestroyWatchlist)

		items := watchlists.Group("/:watchlist_id/items")
		{
			items.GET("", handlers.GetItemsWatchlist)
			items.POST("", handlers.CreateItemWatchlist)
			items.PATCH("/:item_id", handlers.UpdateItemWatchlist)
			items.DELETE("/:item_id", handlers.DestroyItemWatchlist)
		}
	}
}
