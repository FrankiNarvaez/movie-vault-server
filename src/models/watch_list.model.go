package models

type WatchList struct {
	ID     *string `json:"id" db:"id"`
	UserID string  `json:"user_id" db:"user_id"`
	ImdbID string  `json:"imdb_id" db:"imdb_id"`
	Type   string  `json:"type" db:"type"`
}
