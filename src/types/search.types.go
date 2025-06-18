package types

import "movie/src/models"

type MovieResult models.Movie

type SearchMovies struct {
	Page         int           `json:"page"`
	Results      []MovieResult `json:"results"`
	TotalResults int           `json:"total_results"`
	TotalPages   int           `json:"total_pages"`
}

type SearchCelebrities struct {
	Page         int                `json:"page"`
	Results      []models.Celebrity `json:"results"`
	TotalResults int                `json:"total_results"`
	TotalPages   int                `json:"total_pages"`
}

type SearchCompanies struct {
	Page         int              `json:"page"`
	Results      []models.Company `json:"results"`
	TotalResults int              `json:"total_results"`
	TotalPages   int              `json:"total_pages"`
}

type SearchSeries struct {
	Page         int             `json:"page"`
	Results      []ResultsSeries `json:"results"`
	TotalResults int             `json:"total_results"`
	TotalPages   int             `json:"total_pages"`
}

type ResultsAll struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	Id               int     `json:"id"`
	Title            string  `json:"title,omitempty"`
	Name             string  `json:"name"`
	OriginalName     string  `json:"original_name,omitempty"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	MediaType        string  `json:"media_type"`
	GenreIds         []int   `json:"genre_ids"`
	Popularity       float64 `json:"popularity"`
	ReleaseDate      string  `json:"release_date"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type SearchAll struct {
	Page         int          `json:"page"`
	Results      []ResultsAll `json:"results"`
	TotalResults int          `json:"total_results"`
	TotalPages   int          `json:"total_pages"`
}

type SearchResult struct {
	ID               int     `json:"id"`
	Title            string  `json:"title,omitempty"`
	Name             string  `json:"name,omitempty"`
	Overview         string  `json:"overview,omitempty"`
	MediaType        string  `json:"media_type"`
	VoteAverage      float64 `json:"vote_average"`
	OriginalLanguage string  `json:"original_language"`
	GenreIDs         []int   `json:"genre_ids,omitempty"`
	PosterPath       string  `json:"poster_path,omitempty"`
	BackdropPath     string  `json:"backdrop_path,omitempty"`
	Popularity       float64 `json:"popularity"`
}
