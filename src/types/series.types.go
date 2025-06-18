package types

import "movie/src/models"

type ResultsPopularSeries struct {
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

type PopularSeries struct {
	Pages        int64                  `json:"pages"`
	Results      []ResultsPopularSeries `json:"results"`
	TotalPages   int64                  `json:"total_pages"`
	TotalResults int64                  `json:"total_results"`
}

// used to trending and recommendations
type ResultsSeries struct {
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

type TrendingSeries struct {
	Pages        int64           `json:"pages"`
	Results      []ResultsSeries `json:"results"`
	TotalPages   int64           `json:"total_pages"`
	TotalResults int64           `json:"total_results"`
}

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SeriesDetails struct {
	Adult               bool                    `json:"adult"`
	BackdropPath        string                  `json:"backdrop_path"`
	CreatedBy           []models.Creator        `json:"created_by"`
	EpisodeRunTime      []int                   `json:"episode_run_time"`
	FirstAirDate        string                  `json:"first_air_date"`
	Genres              []Genres                `json:"genres"`
	Homepage            string                  `json:"homepage"`
	ID                  int                     `json:"id"`
	InProduction        bool                    `json:"in_production"`
	Languages           []string                `json:"languages"`
	LastAirDate         string                  `json:"last_air_date"`
	LastEpisodeToAir    *models.EpisodeSerie    `json:"last_episode_to_air"`
	Name                string                  `json:"name"`
	NextEpisodeToAir    *models.EpisodeSerie    `json:"next_episode_to_air"`
	Networks            []models.Network        `json:"networks"`
	NumberOfEpisodes    int                     `json:"number_of_episodes"`
	NumberOfSeasons     int                     `json:"number_of_seasons"`
	OriginCountry       []string                `json:"origin_country"`
	OriginalLanguage    string                  `json:"original_language"`
	OriginalName        string                  `json:"original_name"`
	Overview            string                  `json:"overview"`
	Popularity          float64                 `json:"popularity"`
	PosterPath          string                  `json:"poster_path"`
	ProductionCompanies []models.Company        `json:"production_companies"`
	ProductionCountries []models.Country        `json:"production_countries"`
	Seasons             []models.SeasonSerie    `json:"seasons"`
	SpokenLanguages     []models.SpokenLanguage `json:"spoken_languages"`
	Status              string                  `json:"status"`
	Tagline             *string                 `json:"tagline"`
	Type                string                  `json:"type"`
	VoteAverage         float64                 `json:"vote_average"`
	VoteCount           int                     `json:"vote_count"`
}

type CastSeries struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

type SeriesCredits struct {
	Cast []CastSeries        `json:"cast"`
	Crew []models.CrewMember `json:"crew"`
	Id   int                 `json:"id"`
}

// user to series and season
type SeriesImages struct {
	Backdrops *[]Image `json:"backdrops"`
	Id        int      `json:"id"`
	Logos     *[]Image `json:"logos"`
	Posters   []Image  `json:"posters"`
}

// used to unmarshal series and season
type SeriesVideos struct {
	Id      int     `json:"id"`
	Results []Video `json:"results"`
}

type SeriesRecommendations struct {
	Page         int64           `json:"page"`
	Results      []ResultsSeries `json:"results"`
	TotalPages   int64           `json:"total_pages"`
	TotalResults int64           `json:"total_results"`
}

/// Season types

type SeasonDetails struct {
	Id           string           `json:"_id"`
	AirDate      string           `json:"air_date"`
	Episodes     []models.Episode `json:"episodes"`
	Name         string           `json:"name"`
	Overview     string           `json:"overview"`
	ID           int              `json:"id"`
	PosterPath   string           `json:"poster_path"`
	SeasonNumber int              `json:"season_number"`
	VoteAverage  float64          `json:"vote_average"`
}

type SeasonCredits struct {
	Cast []CastSeries        `json:"cast"`
	Crew []models.CrewMember `json:"crew"`
	Id   int                 `json:"id"`
}

// Episodes

type EpisodeCredits struct {
	Cast       []CastSeries        `json:"cast"`
	Crew       []models.CrewMember `json:"crew"`
	GuestStars []models.GuestStar  `json:"guest_stars"`
	Id         int                 `json:"id"`
}

type EpisodeImages struct {
	Id     int     `json:"id"`
	Stills []Image `json:"stills"`
}
