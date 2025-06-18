package handlers

import (
	"database/sql"
	"fmt"
	config "movie/internal"
	"movie/src/errors"
	"movie/src/models"
	"movie/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func verifyWatchlistOwnership(db *sqlx.DB, watchlistID, userID string) error {
	query := `SELECT 1 FROM watchlists WHERE id=$1 AND user_id=$2`
	var exists int
	if err := db.Get(&exists, query, watchlistID, userID); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("watchlist not fount or access denied")
		}
		return err
	}
	return nil
}

func GetWatchlists(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	var watchlists []models.Watchlist
	db := config.DB

	query := `SELECT * FROM watchlists WHERE user_id=$1`
	if err := db.Select(&watchlists, query, user.ID); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not get watchlists"))
		return
	}

	utils.HandleResponseOK(c, watchlists)
}

func CreateWatchlist(c *gin.Context) {
	var body models.Watchlist
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

	db := config.DB
	var watchlist models.Watchlist
	if err := db.Get(&watchlist, query, user.ID, body.Name); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not create watchlist"))
		return
	}

	utils.HandleResponseOK(c, watchlist)
}

func UpdateWatchlist(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	db := config.DB
	id := c.Param("watchlist_id")
	if err := verifyWatchlistOwnership(db, id, user.ID); err != nil {
		utils.HandleError(c, errors.NewNotFoundError(err.Error()))
		return
	}

	var body models.Watchlist
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Invalid request body"))
		return
	}

	queryUpdate := `UPDATE watchlists SET name=$1 WHERE id=$2 AND user_id=$3 RETURNING id, name, user_id`
	var watchlist models.Watchlist
	if err := db.Get(&watchlist, queryUpdate, body.Name, id, user.ID); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not update watchlist"))
		return
	}

	utils.HandleResponseOK(c, watchlist)
}

func DestroyWatchlist(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	db := config.DB
	id := c.Param("watchlist_id")
	if err := verifyWatchlistOwnership(db, id, user.ID); err != nil {
		utils.HandleError(c, errors.NewNotFoundError(err.Error()))
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
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	watchlistID := c.Param("watchlist_id")
	db := config.DB
	if err := verifyWatchlistOwnership(db, watchlistID, user.ID); err != nil {
		utils.HandleError(c, errors.NewNotFoundError(err.Error()))
		return
	}

	var watchlist []models.ItemWatchlist
	query := `SELECT * FROM item_watchlists WHERE watchlist_id=$1`
	if err := db.Select(&watchlist, query, watchlistID); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not get items"))
		return
	}

	utils.HandleResponseOK(c, watchlist)
}

func CreateItemWatchlist(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	db := config.DB
	watchlistID := c.Param("watchlist_id")

	if err := verifyWatchlistOwnership(db, watchlistID, user.ID); err != nil {
		utils.HandleError(c, errors.NewNotFoundError(err.Error()))
		return
	}

	var body models.ItemWatchlist
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Invalid request body"))
		return
	}

	query := `INSERT INTO item_watchlists (watchlist_id, tmdb_id, type) 
          VALUES ($1, $2, $3) 
          RETURNING id, watchlist_id, tmdb_id, type`
	var item models.ItemWatchlist
	if err := db.Get(&item, query, watchlistID, body.TmdbID, body.Type); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not add to watchlist"))
		return
	}

	utils.HandleResponseOK(c, item)
}

func UpdateItemWatchlist(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	db := config.DB
	watchlistID := c.Param("watchlist_id")
	itemID := c.Param("item_id")

	if err := verifyWatchlistOwnership(db, watchlistID, user.ID); err != nil {
		utils.HandleError(c, errors.NewNotFoundError(err.Error()))
		return
	}

	var exists int
	querySearch := `SELECT 1 FROM item_watchlists WHERE id=$1 AND watchlist_id=$2`
	if err := db.Get(&exists, querySearch, itemID, watchlistID); err != nil {
		utils.HandleError(c, errors.NewNotFoundError("Watchlist not found"))
		return
	}

	var body models.ItemWatchlist
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Invalid request body"))
		return
	}

	queryUpdate := `UPDATE item_watchlists SET status=$1 
	WHERE id=$2 AND watchlist_id=$3
	RETURNING id, watchlist_id, type, tmdb_id, status`
	var item models.ItemWatchlist
	if err := db.Get(&item, queryUpdate, body.Status, itemID, watchlistID); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not update watchlist"))
		return
	}

	utils.HandleResponseOK(c, item)
}

func DestroyItemWatchlist(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		utils.HandleError(c, errors.NewUnauthorizedError(err.Error()))
		return
	}

	db := config.DB
	watchlistID := c.Param("watchlist_id")
	itemID := c.Param("item_id")

	if err := verifyWatchlistOwnership(db, watchlistID, user.ID); err != nil {
		utils.HandleError(c, errors.NewNotFoundError(err.Error()))
		return
	}

	querySearch := `SELECT 1 FROM item_watchlists WHERE watchlist_id=$1 AND id=$2`
	var exists int
	if err := db.Get(&exists, querySearch, watchlistID, itemID); err != nil {
		utils.HandleError(c, errors.NewNotFoundError("Item from watchlist not found"))
		return
	}

	query_destroy := `DELETE FROM item_watchlists WHERE watchlist_id=$1 AND id=$2`
	if _, err := db.Exec(query_destroy, watchlistID, itemID); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not remove from watchlist"))
		return
	}

	utils.HandleResponseOK(c, nil)
}
