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
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	query := `SELECT id, user_id, tmdb_id, type from favorites WHERE user_id=$1`

	if err := db.Select(&favorites, query, user.ID); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not get favorites"))
		return
	}

	utils.HandleResponseOK(c, favorites)
}

func CreateFavorite(c *gin.Context) {
	var body models.Favorite
	db := config.DB
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Invalid request body"))
		return
	}

	query := `INSERT INTO favorites (user_id, tmdb_id, type)
	VALUES ($1, $2, $3)
	RETURNING id, user_id, tmdb_id, type`

	var favorite models.Favorite
	if err := db.Get(&favorite, query, user.ID, body.TmdbID, body.Type); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not add to favorites"))
		return
	}

	utils.HandleResponseOK(c, favorite)
}

func DestroyFavorite(c *gin.Context) {
	db := config.DB
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	id := c.Param("id")
	query_search := `SELECT 1 FROM favorites WHERE id=$1 and user_id=$2`
	var exists int

	if err := db.Get(&exists, query_search, id, user.ID); err != nil {
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
