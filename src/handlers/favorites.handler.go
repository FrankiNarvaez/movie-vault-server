package handlers

import (
	config "movie/internal"
	"movie/src/errors"
	"movie/src/models"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	var favorites []models.Favorite
	db := config.DB

	query := `SELECT id, user_id, tmdb_id, type from favorites`

	if err := db.Select(&favorites, query); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not get favorites"))
		return
	}

	utils.HandleResponseOK(c, favorites)
}

func CreateFavorite(c *gin.Context) {
	var favorite models.Favorite
	db := config.DB

	if err := c.ShouldBindJSON(&favorite); err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Invalid request body"))
		return
	}

	query := `INSERT INTO favorites (user_id, tmdb_id, type) VALUES ($1, $2, $3)`

	if _, err := db.Query(query, favorite.UserID, favorite.TmdbID, favorite.Type); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not add to favorites"))
		return
	}

	utils.HandleResponseOK(c, favorite)
}

func DestroyFavorite(c *gin.Context) {
	db := config.DB

	id := c.Param("id")
	query_search := `SELECT 1 FROM favorites WHERE id=$1`
	var exists int

	if err := db.Get(&exists, query_search, id); err != nil {
		utils.HandleError(c, errors.NewNotFoundError("Favorite not found"))
		return
	}

	query_destroy := `DELETE FROM favorites WHERE id=$1`

	if _, err := db.Exec(query_destroy, id); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not remove from favorites"))
		return
	}

	utils.HandleResponseOK(c, nil)
}
