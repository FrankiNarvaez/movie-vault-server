package types

import "movie/src/models"

type Movies struct {
	Page         int64          `json:"page"`
	Results      []models.Movie `json:"results"`
	TotalPages   int64          `json:"total_pages"`
	TotalResults int64          `json:"total_results"`
}

type VideoResults struct {
	ID      int            `json:"id"`
	Results []models.Video `json:"results"`
}
