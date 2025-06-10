package models

type Movie struct {
	Adult               bool                `json:"adult"`
	BackdropPath        string              `json:"backdrop_path"`
	BelongsToCollection *Collection         `json:"belongs_to_collection"`
	Budget              int                 `json:"budget"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ID                  int                 `json:"id"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalTitle       string              `json:"original_title"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	ReleaseDate         string              `json:"release_date"`
	Revenue             int64               `json:"revenue"`
	Runtime             int                 `json:"runtime"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Title               string              `json:"title"`
	Video               bool                `json:"video"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`
}

type Collection struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductionCompany struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

type ProductionCountry struct {
	ISO3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

type SpokenLanguage struct {
	EnglishName string `json:"english_name"`
	ISO639_1    string `json:"iso_639_1"`
	Name        string `json:"name"`
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
	Iso639_1    string  `json:"iso_639_1"`
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

type ExternalIDs struct {
	ID          int    `json:"id"`
	ImdbID      string `json:"imdb_id"`
	WikidataID  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
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
