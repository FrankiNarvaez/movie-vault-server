package types

type DiscorverSeriesResults struct {
	Pages        int64           `json:"pages"`
	Results      []DiscoverSerie `json:"results"`
	TotalPages   int64           `json:"total_pages"`
	TotalResults int64           `json:"total_results"`
}

type DiscoverSerie struct {
	BackdropPath     string   `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

type DiscorverMoviesResults struct {
	Pages        int64           `json:"pages"`
	Results      []DiscoverMovie `json:"results"`
	TotalPages   int64           `json:"total_pages"`
	TotalResults int64           `json:"total_results"`
}

type DiscoverMovie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	FirstAirDate     string  `json:"first_air_date"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_tile"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}
