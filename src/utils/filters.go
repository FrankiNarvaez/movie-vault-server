package utils

import (
	"movie/src/models"
	"movie/src/types"
	"slices"
)

func ApplyFilters(results []types.SearchResult, filters models.Filters) []types.SearchResult {
	var filtered []types.SearchResult

	for _, item := range results {
		if filters.MediaType != nil && item.MediaType != *filters.MediaType {
			continue
		}

		// Filtrar por votos
		if filters.MinVote != nil && *filters.MinVote > 0 && item.VoteAverage < *filters.MinVote {
			continue
		}
		if filters.MaxVote != nil && *filters.MaxVote > 0 && item.VoteAverage > *filters.MaxVote {
			continue
		}

		// Filtrar por géneros
		if len(*filters.GenreIDs) > 0 && !hasAnyGenre(item.GenreIDs, *filters.GenreIDs) {
			continue
		}

		filtered = append(filtered, item)
	}

	return filtered
}

func hasAnyGenre(genres []int, filters []int) bool {
	for _, g := range genres {
		if slices.Contains(filters, g) {
			return true
		}
	}
	return false
}

func IsFiltersEmpty(f models.Filters) bool {
	return (f.MediaType == nil || *f.MediaType == "") &&
		(f.MinVote == nil || *f.MinVote == 0) &&
		(f.MaxVote == nil || *f.MaxVote == 0) &&
		(f.GenreIDs == nil || len(*f.GenreIDs) == 0)
}

func PassesFilters(m models.Movie, filters models.Filters) bool {
	// Filtrar por votos
	if filters.MinVote != nil && m.VoteAverage < *filters.MinVote {
		return false
	}
	if filters.MaxVote != nil && m.VoteAverage > *filters.MaxVote {
		return false
	}

	if len(*filters.GenreIDs) > 0 && !hasAnyGenre(*m.Genres_IDs, *filters.GenreIDs) {
		return false
	}

	return true
}
