package models

type Filters struct {
	MediaType *string  `json:"media_type"`
	GenreIDs  *[]int   `json:"genre_ids"`
	MinVote   *float64 `json:"min_vote"`
	MaxVote   *float64 `json:"max_vote"`
}
