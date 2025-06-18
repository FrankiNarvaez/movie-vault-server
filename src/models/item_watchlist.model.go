package models

import "time"

type ItemWatchlist struct {
	ID          string    `json:"id" db:"id"`
	WatchlistID string    `json:"watchlist_id" db:"watchlist_id"`
	TmdbID      string    `json:"tmdb_id" db:"tmdb_id"`
	Type        string    `json:"type" db:"type"`
	Status      int       `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
