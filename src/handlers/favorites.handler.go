package handlers

import (
	config "movie/internal"
	"movie/src/errors"
	"movie/src/models"
	"movie/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	var favorites []models.Favorite
	db := config.DB

	query := `SELECT id, user_id, imdb_id, type from favorites`

	if err := db.Select(&favorites, query); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not get favorites"))
		return
	}

	c.JSON(http.StatusOK, favorites)
}

func CreateFavorite(c *gin.Context) {
	var favorite models.Favorite
	db := config.DB

	if err := c.ShouldBindJSON(&favorite); err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Invalid request body"))
		return
	}

	favorite.ImdbID = c.Param("imdb_id")

	query := `INSERT INTO favorites (user_id, imdb_id, type) VALUES ($1, $2, $3)`

	if _, err := db.Query(query, favorite.UserID, favorite.ImdbID, favorite.Type); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not add to favorites"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully added to favorites",
		"data":    favorite,
	})
}

func DestroyFavorite(c *gin.Context) {
	var favorite models.Favorite
	db := config.DB

	if err := c.ShouldBindJSON(&favorite); err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Invalid request body"))
		return
	}

	favorite.ImdbID = c.Param("imdb_id")

	query_search := `SELECT 1 FROM favorites WHERE user_id = $1 AND imdb_id = $2 AND type = $3`
	var exists int

	if err := db.Get(&exists, query_search, favorite.UserID, favorite.ImdbID, favorite.Type); err != nil {
		utils.HandleError(c, errors.NewNotFoundError("Favorite not found"))
		return
	}

	query_destroy := `DELETE FROM favorites WHERE user_id = $1 AND imdb_id = $2 AND type = $3;`

	if _, err := db.Exec(query_destroy, favorite.UserID, favorite.ImdbID, favorite.Type); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not remove from favorites"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully removed from favorites"})
}
