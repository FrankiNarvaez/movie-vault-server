package routes

import (
	"movie/src/handlers"

	"github.com/gin-gonic/gin"
)

func WatchListRoutes(api *gin.RouterGroup) {
	api.GET("/watch_list", handlers.GetFromWatchLists)
	api.POST("/watch_list/:imdb_id", handlers.CreateAtWatchList)
	api.DELETE("/watch_list/:imdb_id", handlers.DeleteFromWatchList)
}
