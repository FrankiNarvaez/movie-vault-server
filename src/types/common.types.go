package types

type ExternalIDs struct {
	Id           int     `json:"id"`
	Freebase_mid *string `json:"freebase_mid,omitempty"`
	Freebase_sid *string `json:"freebase_sid,omitempty"`
	Freebase_id  *string `json:"freebase_id,omitempty"`
	Tmdb_id      *string `json:"tmdb_id,omitempty"`
	Serierage_id *int64  `json:"tvrage_id,omitempty"`
	SeriedbId    *int64  `json:"tvdb_id,omitempty"`
	Wikidata_id  *string `json:"wikidata_id,omitempty"`
	Facebook_id  *string `json:"facebook_id,omitempty"`
	Instagram_id *string `json:"instagram_id,omitempty"`
	Tiktok_id    *string `json:"tiktok_id,omitempty"`
	Twitter_id   *string `json:"twitter_id,omitempty"`
	Youtube_id   *string `json:"youtube_id,omitempty"`
}

type WatchProviders struct {
	ID      int                        `json:"id"`
	Results map[string]CountryProvider `json:"results"`
}

type CountryProvider struct {
	Link     string     `json:"link"`
	Flatrate []Provider `json:"flatrate"`
	Rent     []Provider `json:"rent"`
	Buy      []Provider `json:"buy"`
}

type Provider struct {
	LogoPath        string `json:"logo_path"`
	ProviderID      int    `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
	DisplayPriority int    `json:"display_priority"`
}

type TrendingResults struct {
	Pages        int64            `json:"pages"`
	Results      []TrendingResult `json:"results"`
	TotalPages   int64            `json:"total_pages"`
	TotalResults int64            `json:"total_results"`
}

type TrendingResult struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	MediaType        string   `json:"media_type"`
	GenreIds         []int    `json:"genre_ids"`
	Popularity       float64  `json:"popularity"`
	FirstAirDate     string   `json:"first_air_date"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	OriginCountry    []string `json:"origin_country"`
}
