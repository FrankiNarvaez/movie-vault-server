package models

// TV
type TVShow struct {
	Adult               bool       `json:"adult"`
	BackdropPath        string     `json:"backdrop_path"`
	CreatedBy           []Creator  `json:"created_by"`
	EpisodeRunTime      []int      `json:"episode_run_time"`
	FirstAirDate        string     `json:"first_air_date"`
	Genres              []Genre    `json:"genres"`
	Homepage            string     `json:"homepage"`
	ID                  int        `json:"id"`
	InProduction        bool       `json:"in_production"`
	Languages           []string   `json:"languages"`
	LastAirDate         string     `json:"last_air_date"`
	LastEpisodeToAir    *Episode   `json:"last_episode_to_air"`
	Name                string     `json:"name"`
	NextEpisodeToAir    *EpisodeTv `json:"next_episode_to_air"`
	Networks            []Network  `json:"networks"`
	NumberOfEpisodes    int        `json:"number_of_episodes"`
	NumberOfSeasons     int        `json:"number_of_seasons"`
	OriginCountry       []string   `json:"origin_country"`
	OriginalLanguage    string     `json:"original_language"`
	OriginalName        string     `json:"original_name"`
	Overview            string     `json:"overview"`
	Popularity          float64    `json:"popularity"`
	PosterPath          string     `json:"poster_path"`
	ProductionCompanies []Company  `json:"production_companies"`
	ProductionCountries []Country  `json:"production_countries"`
	Seasons             []SeasonTV `json:"seasons"`
}

type Creator struct {
	ID           int     `json:"id"`
	CreditID     string  `json:"credit_id"`
	Name         string  `json:"name"`
	OriginalName string  `json:"original_name"`
	Gender       int     `json:"gender"`
	ProfilePath  *string `json:"profile_path"`
}

type EpisodeTv struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	EpisodeType    *string `json:"episode_type"`
	ProductionCode *string `json:"production_code"`
	Runtime        *int    `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
}

type Network struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

type Company struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

type Country struct {
	ISO3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

type SeasonTV struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     *string `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float64 `json:"vote_average"`
}

// Seasons
type SeasonSeason struct {
	ID       string    `json:"_id"`
	AirDate  string    `json:"air_date"`
	Episodes []Episode `json:"episodes"`
}

// duplicated -- season - episode
type Episode struct {
	AirDate        string       `json:"air_date"`
	EpisodeNumber  int          `json:"episode_number"`
	EpisodeType    *string      `json:"episode_type"`
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Overview       string       `json:"overview"`
	ProductionCode string       `json:"production_code"`
	Runtime        int          `json:"runtime"`
	SeasonNumber   int          `json:"season_number"`
	ShowID         *int         `json:"show_id"`
	StillPath      string       `json:"still_path"`
	VoteAverage    float64      `json:"vote_average"`
	VoteCount      int          `json:"vote_count"`
	Crew           []CrewMember `json:"crew"`
	GuestStars     []GuestStar  `json:"guest_stars"`
}

// duplicated -- season - episode
type CrewMember struct {
	Job                string  `json:"job"`
	Department         string  `json:"department"`
	CreditID           string  `json:"credit_id"`
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

// duplicated -- season - episode
type GuestStar struct {
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Order              int     `json:"order"`
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
}

// Episodes
type EpisodeEpisode struct {
	AirDate        string       `json:"air_date"`
	Crew           []CrewMember `json:"crew"`
	EpisodeNumber  int          `json:"episode_number"`
	EpisodeType    string       `json:"episode_type"`
	GuestStars     []GuestStar  `json:"guest_stars"`
	Name           string       `json:"name"`
	Overview       string       `json:"overview"`
	ID             int          `json:"id"`
	ProductionCode string       `json:"production_code"`
	Runtime        int          `json:"runtime"`
	SeasonNumber   int          `json:"season_number"`
	StillPath      string       `json:"still_path"`
	VoteAverage    float64      `json:"vote_average"`
	VoteCount      int          `json:"vote_count"`
}
