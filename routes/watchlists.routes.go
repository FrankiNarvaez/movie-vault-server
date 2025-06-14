package routes

import (
	"movie/src/handlers"
	"movie/src/middlewares"

	"github.com/gin-gonic/gin"
)

func WatchListRoutes(api *gin.RouterGroup) {
	api.GET("/watchlist", middlewares.CheckAuth, handlers.GetFromWatchLists)
	api.POST("/watchlist/:imdb_id", middlewares.CheckAuth, handlers.CreateAtWatchList)
	api.DELETE("/watchlist/:imdb_id", middlewares.CheckAuth, handlers.DeleteFromWatchList)
}
