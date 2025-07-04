package types

import (
	"movie/src/models"
)

// Define a local type alias for models.Movie
type (
	LocalMovie models.Movie
	LocalSerie ResultsSeries
)

func (s SearchResult) ToMovieResult() LocalMovie {
	return LocalMovie{
		ID:               s.ID,
		Title:            s.Title,
		VoteAverage:      s.VoteAverage,
		OriginalLanguage: s.OriginalLanguage,
		Genres_IDs:       &s.GenreIDs,
	}
}

// Para SerieResult (si lo usas)
func (t LocalSerie) ToSearchResult() SearchResult {
	return SearchResult{
		ID:               t.ID,
		Name:             t.Name,
		MediaType:        "serie",
		VoteAverage:      t.VoteAverage,
		OriginalLanguage: t.OriginalLanguage,
		GenreIDs:         t.GenreIds,
	}
}

func (s SearchResult) ToSerieResult() LocalSerie {
	return LocalSerie{
		ID:               s.ID,
		Name:             s.Name,
		VoteAverage:      s.VoteAverage,
		OriginalLanguage: s.OriginalLanguage,
		GenreIds:         s.GenreIDs,
	}
}
