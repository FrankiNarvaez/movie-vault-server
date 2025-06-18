package types

import "movie/src/models"

type CelebrityTrending struct {
	Adult              bool    `json:"adult"`
	Id                 int     `json:"id"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float32 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	MediaType          string  `json:"media_type"`
	KnownForDepartment string  `json:"known_for_department"`
	Gender             int     `json:"gender"`
}

// used to unmarshal images
type Profiles struct {
	AspectRatio  float64 `json:"aspect_ratio"`
	Height       int     `json:"height"`
	Iso_639_1    *string `json:"iso_639_1"`
	File_path    string  `json:"file_path"`
	Vote_average float64 `json:"vote_average"`
	Vote_count   int     `json:"vote_count"`
	Width        int     `json:"width"`
}

// used to unmarshal credits and combined credits
type Cast struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        float64 `json:"vote_count"`
	Character        string  `json:"character"`
	CreditId         string  `json:"credit_id"`
	Order            int64   `json:"order"`
	MediaType        *string `json:"media_type"`
}

type Crew struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        float64 `json:"vote_count"`
	CreditId         string  `json:"credit_id"`
	Department       string  `json:"department"`
	Job              string  `json:"job"`
	MediaType        *string `json:"media_type"`
}

type CelebrityExternalIds struct {
	Id           int     `json:"id"`
	Freebase_mid *string `json:"freebase_mid"`
	Freebase_sid *string `json:"freebase_sid"`
	Tmdb_id      *string `json:"tmdb_id"`
	Serierage_id *int64  `json:"tvrage_id"`
	Wikidata_id  *string `json:"wikidata_id"`
	Facebook_id  *string `json:"facebook_id"`
	Instagram_id *string `json:"instagram_id"`
	Tiktok_id    *string `json:"tiktok_id"`
	Twitter_id   *string `json:"twitter_id"`
	Youtube_id   *string `json:"youtube_id"`
}

type PopularCelebrities struct {
	Pages        int64              `json:"pages"`
	Results      []models.Celebrity `json:"results"`
	TotalPages   int64              `json:"total_pages"`
	TotalResults int64              `json:"total_results"`
}

type TrendingCelebrities struct {
	Pages        int64               `json:"pages"`
	Results      []CelebrityTrending `json:"results"`
	TotalPages   int64               `json:"total_pages"`
	TotalResults int64               `json:"total_results"`
}

type CelebrityImages struct {
	Id       int        `json:"id"`
	Profiles []Profiles `json:"profiles"`
}

type CelebrityMovieCredits struct {
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
	Id   int    `json:"id"`
}

type CelebrityCombinedCredits struct {
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
	Id   int    `json:"id"`
}
