package routes

import (
	"movie/src/handlers"
	"movie/src/middlewares"

	"github.com/gin-gonic/gin"
)

func WatchListRoutes(api *gin.RouterGroup) {
	watchlists := api.Group("watchlists")
	{
		watchlists.GET("", middlewares.CheckAuth, handlers.GetWatchlists)
		watchlists.POST("/:watchlist_id", middlewares.CheckAuth, handlers.CreateWatchlist)
		watchlists.DELETE("/:watchlist_id", middlewares.CheckAuth, handlers.DestroyWatchlist)

		items := watchlists.Group("/:watchlist_id/items")
		{
			items.GET("", handlers.GetItemsWatchlist)
			items.POST("/:item_id", handlers.CreateItemWatchlist)
			items.PATCH("/:item_id", handlers.UpdateItemWatchlist)
			items.DELETE("/:item_id", handlers.DestroyItemWatchlist)
		}
	}
}
