package handlers

import (
	config "movie/internal"
	"movie/src/errors"
	"movie/src/models"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
)

func GetWatchlists(c *gin.Context) {
	var watchlists []models.Watchlist
	db := config.DB
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	query := `SELECT * FROM watchlists WHERE user_id=$1`

	if err := db.Select(&watchlists, query, user.ID); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not get watchlists"))
		return
	}

	utils.HandleResponseOK(c, watchlists)
}

func CreateWatchlist(c *gin.Context) {
	var body models.Watchlist
	db := config.DB
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Invalid request body"))
		return
	}

	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	query := `INSERT INTO watchlists (user_id, name)
	VALUES ($1, $2)
	RETURNING id, user_id, name`

	var watchlist models.Watchlist
	if err := db.Get(&watchlist, query, user.ID, body.Name); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not create watchlist"))
		return
	}

	utils.HandleResponseOK(c, watchlist)
}

func DestroyWatchlist(c *gin.Context) {
	db := config.DB
	id := c.Param("watchlist_id")

	_, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	querySearch := `SELECT 1 FROM watchlists WHERE id=$1`

	var exists int
	if err := db.Get(&exists, querySearch, id); err != nil {
		utils.HandleError(c, errors.NewNotFoundError("Watchlist not found"))
		return
	}

	queryDelete := `DELETE FROM watchlists WHERE id=$1`
	if _, err := db.Exec(queryDelete, id); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not remove from watchlist"))
		return
	}

	utils.HandleResponseOK(c, nil)
}

func GetItemsWatchlist(c *gin.Context) {
	var watchlist []models.ItemWatchlist
	db := config.DB
	_, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	watchlist_id := c.Param("watchlist_id")

	query := `SELECT id, watchlist_id, tmdb_id, type from item_watchlists WHERE watchlist_id=$1`

	if err := db.Select(&watchlist, query, watchlist_id); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not get items"))
		return
	}

	utils.HandleResponseOK(c, watchlist)
}

func CreateItemWatchlist(c *gin.Context) {
	var body models.ItemWatchlist
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

	query := `INSERT INTO item_watchlists (watchlist_id, tmdb_id, type) VALUES ($1, $2, $3)`

	var item models.ItemWatchlist
	if err := db.Get(&item, query, user.ID, body.TmdbID); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not add to watchlist"))
		return
	}

	utils.HandleResponseOK(c, item)
}

func UpdateItemWatchlist(c *gin.Context) {}

func DestroyItemWatchlist(c *gin.Context) {
	db := config.DB
	watchlist_id := c.Param("watchlist_id")
	item_id := c.Param("item_id")

	_, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}
	querySearch := `SELECT 1 FROM item_watchlists WHERE watchlist_id=$1 AND id=$2`
	var exists int

	if err := db.Get(&exists, querySearch, watchlist_id, item_id); err != nil {
		utils.HandleError(c, errors.NewNotFoundError("Item from watchlist not found"))
		return
	}

	query_destroy := `DELETE FROM item_watchlists WHERE watchlist_id=$1 AND id=$2`

	if _, err := db.Exec(query_destroy, watchlist_id, item_id); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not remove from watchlist"))
		return
	}

	utils.HandleResponseOK(c, nil)
}
