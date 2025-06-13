package types

import "movie/src/models"

type Movies struct {
	Page         int64          `json:"page"`
	Results      []models.Movie `json:"results"`
	TotalPages   int64          `json:"total_pages"`
	TotalResults int64          `json:"total_results"`
}

type VideoResults struct {
	ID      int     `json:"id"`
	Results []Video `json:"results"`
}

type MovieCredits struct {
	ID   int                 `json:"id"`
	Cast []CastMember        `json:"cast"`
	Crew []CrewMemberCredits `json:"crew"`
}

type CastMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CastID             int     `json:"cast_id"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

type CrewMemberCredits struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CreditID           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Job                string  `json:"job"`
}

type MediaImages struct {
	ID        int     `json:"id"`
	Backdrops []Image `json:"backdrops"`
	Logos     []Image `json:"logos"`
	Posters   []Image `json:"posters"`
}

type Image struct {
	AspectRatio float64 `json:"aspect_ratio"`
	Height      int     `json:"height"`
	Iso639_1    *string `json:"iso_639_1"`
	FilePath    string  `json:"file_path"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}

type Video struct {
	Iso639_1    string `json:"iso_639_1"`
	Iso3166_1   string `json:"iso_3166_1"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Site        string `json:"site"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
	Official    bool   `json:"official"`
	PublishedAt string `json:"published_at"`
	ID          string `json:"id"`
}
