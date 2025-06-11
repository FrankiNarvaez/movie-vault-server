package handlers

import (
	config "movie/internal"
	"movie/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFromWatchLists(c *gin.Context) {
	var watch_lists []models.WatchList
	db := config.DB

	query := `SELECT id, user_id, imdb_id, type from watch_lists`

	if err := db.Select(&watch_lists, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get favorites"})
		return
	}

	c.JSON(http.StatusOK, watch_lists)
}

func CreateAtWatchList(c *gin.Context) {
	var watch_list models.WatchList
	db := config.DB

	if err := c.ShouldBindJSON(&watch_list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	watch_list.ImdbID = c.Param("imdb_id")

	query := `INSERT INTO watch_lists (user_id, imdb_id, type) VALUES ($1, $2, $3)`

	if _, err := db.Query(query, watch_list.UserID, watch_list.ImdbID, watch_list.Type); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add to favorites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully added to favorites",
		"data":    watch_list,
	})
}

func DeleteFromWatchList(c *gin.Context) {
	var watch_list models.WatchList
	db := config.DB

	if err := c.ShouldBindJSON(&watch_list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	watch_list.ImdbID = c.Param("imdb_id")

	query_search := `SELECT 1 FROM watch_lists WHERE user_id = $1 AND imdb_id = $2 AND type = $3`
	var exists int

	if err := db.Get(&exists, query_search, watch_list.UserID, watch_list.ImdbID, watch_list.Type); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
		return
	}

	query_destroy := `DELETE FROM watch_lists WHERE user_id = $1 AND imdb_id = $2 AND type = $3;`

	if _, err := db.Exec(query_destroy, watch_list.UserID, watch_list.ImdbID, watch_list.Type); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove from favorites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully removed from favorites"})
}
