package utils

import (
	"movie/src/models"
	"movie/src/types"
	"slices"
	"strings"
)

func ApplyFilters(results []types.SearchResult, filters models.Filters) []types.SearchResult {
	var filtered []types.SearchResult

	for _, item := range results {
		if filters.MediaType != "" && item.MediaType != filters.MediaType {
			continue
		}

		if len(filters.Language) > 0 && !contains(filters.Language, item.OriginalLanguage) {
			continue
		}

		// Filtrar por votos
		if filters.MinVote > 0 && item.VoteAverage < filters.MinVote {
			continue
		}
		if filters.MaxVote > 0 && item.VoteAverage > filters.MaxVote {
			continue
		}

		// Filtrar por géneros (si aplica)
		if len(filters.GenreIDs) > 0 && !hasAnyGenre(item.GenreIDs, filters.GenreIDs) {
			continue
		}

		filtered = append(filtered, item)
	}

	return filtered
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.ToLower(item) == strings.ToLower(val) {
			return true
		}
	}
	return false
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
	return f.MediaType == "" && len(f.Language) == 0 && f.MinVote == 0 && f.MaxVote == 0 && len(f.GenreIDs) == 0
}

func PassesFilters(m models.Movie, filters models.Filters) bool {

	if len(filters.Language) > 0 && !contains(filters.Language, m.OriginalLanguage) {
		return false
	}

	// Filtrar por votos
	if filters.MinVote > 0 && m.VoteAverage < filters.MinVote {
		return false
	}
	if filters.MaxVote > 0 && m.VoteAverage > filters.MaxVote {
		return false
	}

	// Filtrar por géneros (si aplica)
	if len(filters.GenreIDs) > 0 && !hasAnyGenre(*m.Genres_IDs, filters.GenreIDs) {
		return false
	}

	return true
}
