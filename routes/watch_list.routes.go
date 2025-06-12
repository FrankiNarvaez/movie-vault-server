package routes

import (
	"movie/src/handlers"
	"movie/src/middlewares"

	"github.com/gin-gonic/gin"
)

func WatchListRoutes(api *gin.RouterGroup) {
	api.GET("/watch_list", middlewares.CheckAuth, handlers.GetFromWatchLists)
	api.POST("/watch_list/:imdb_id", middlewares.CheckAuth, handlers.CreateAtWatchList)
	api.DELETE("/watch_list/:imdb_id", middlewares.CheckAuth, handlers.DeleteFromWatchList)
}
