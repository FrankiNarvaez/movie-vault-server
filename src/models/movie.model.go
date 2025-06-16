package models

type Movie struct {
	Adult               bool                 `json:"adult"`
	BackdropPath        *string              `json:"backdrop_path"`
	BelongsToCollection *Collection          `json:"belongs_to_collection"`
	Budget              *int                 `json:"budget"`
	Genres              *[]Genre             `json:"genres"`
	Genres_IDs          *[]int               `json:"genre_ids"`
	Homepage            *string              `json:"homepage"`
	ID                  int                  `json:"id"`
	OriginCountry       *[]string            `json:"origin_country"`
	OriginalLanguage    string               `json:"original_language"`
	OriginalTitle       string               `json:"original_title"`
	Overview            string               `json:"overview"`
	Popularity          float64              `json:"popularity"`
	PosterPath          string               `json:"poster_path"`
	ProductionCompanies *[]ProductionCompany `json:"production_companies"`
	ProductionCountries *[]ProductionCountry `json:"production_countries"`
	ReleaseDate         string               `json:"release_date"`
	Revenue             *int64               `json:"revenue"`
	Runtime             *int                 `json:"runtime"`
	SpokenLanguages     *[]SpokenLanguage    `json:"spoken_languages"`
	Status              *string              `json:"status"`
	Tagline             *string              `json:"tagline"`
	Title               string               `json:"title"`
	Video               bool                 `json:"video"`
	VoteAverage         float64              `json:"vote_average"`
	VoteCount           int                  `json:"vote_count"`
}

type SearchResult struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	MediaType        string  `json:"media_type"`
	VoteAverage      float64 `json:"vote_average"`
	OriginalLanguage string  `json:"original_language"`
	GenreIDs         []int   `json:"genre_ids"`
}

func (m Movie) ToSearchResult() SearchResult {
	return SearchResult{
		ID:               m.ID,
		Title:            m.Title,
		MediaType:        "movie",
		VoteAverage:      m.VoteAverage,
		OriginalLanguage: m.OriginalLanguage,
		GenreIDs:         derefIntSlice(m.Genres_IDs),
	}
}

// Helper function to safely dereference *[]int
func derefIntSlice(ptr *[]int) []int {
	if ptr != nil {
		return *ptr
	}
	return []int{}
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
