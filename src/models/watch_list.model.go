package models

type WatchList struct {
	ID     *string `json:"id" db:"id"`
	UserID string  `json:"user_id" db:"user_id"`
	TmdbID string  `json:"tmdb_id" db:"tmdb_id"`
	Type   string  `json:"type" db:"type"`
}
