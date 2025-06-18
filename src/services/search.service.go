package services

import (
	"fmt"
	"movie/src/models"
	"movie/src/types"
	"movie/src/utils"
)

func FetchSearchMovies(language, query string, include_adult bool, page string, filters models.Filters) (types.SearchMovies, error) {
	var movies types.SearchMovies

	url := fmt.Sprintf("/search/movie?language=%s&query=%s&include_adult=%t&page=%s", language, query, include_adult, page)

	statusCode, err := utils.FetchFromTMDB(url, &movies)
	if err != nil {
		return types.SearchMovies{}, utils.HandleTMDBError(statusCode, "search movies for query "+query, err)
	}

	if !utils.IsFiltersEmpty(filters) {
		filtered := make([]types.MovieResult, 0, len(movies.Results))
		for _, m := range movies.Results {
			modelMovie := models.Movie{
				Genres_IDs:       m.Genres_IDs,
				VoteAverage:      m.VoteAverage,
				OriginalLanguage: m.OriginalLanguage,
			}
			if utils.PassesFilters(modelMovie, filters) {
				filtered = append(filtered, m)
			}
		}
		movies.Results = filtered
		movies.TotalResults = len(filtered)
	}

	return movies, nil
}

func FetchSearchSeries(language, query string, include_adult bool, page string, filters models.Filters) (types.SearchSeries, error) {
	var series types.SearchSeries

	url := fmt.Sprintf("/search/tv?query=%s&include_adult=%t&language=%s&page=%s", query, include_adult, language, page)

	statusCode, err := utils.FetchFromTMDB(url, &series)
	if err != nil {
		return types.SearchSeries{}, utils.HandleTMDBError(statusCode, "search series for query "+query, err)
	}

	if !utils.IsFiltersEmpty(filters) {
		filtered := make([]types.ResultsSeries, 0, len(series.Results))
		for _, s := range series.Results {
			modelSeries := models.Movie{
				Genres_IDs:       &s.GenreIds,
				VoteAverage:      s.VoteAverage,
				OriginalLanguage: s.OriginalLanguage,
			}
			if utils.PassesFilters(modelSeries, filters) {
				filtered = append(filtered, s)
			}
		}
		series.Results = filtered
		series.TotalResults = len(filtered)
	}

	return series, nil
}

func FetchSearchAll(language, query string, include_adult bool, page string, filters models.Filters) (types.SearchAll, error) {
	var all types.SearchAll

	url := fmt.Sprintf("/search/multi?query=%s&include_adult=%t&language=%s&page=%s", query, include_adult, language, page)

	statusCode, err := utils.FetchFromTMDB(url, &all)
	if err != nil {
		return types.SearchAll{}, utils.HandleTMDBError(statusCode, "search all for query "+query, err)
	}

	if !utils.IsFiltersEmpty(filters) {
		var searchResults []types.SearchResult
		for _, item := range all.Results {
			sr := types.SearchResult{
				ID:               item.Id,
				Title:            item.Title,
				Name:             item.Name,
				MediaType:        item.MediaType,
				Overview:         item.Overview,
				VoteAverage:      item.VoteAverage,
				OriginalLanguage: item.OriginalLanguage,
				GenreIDs:         item.GenreIds,
				PosterPath:       item.PosterPath,
				BackdropPath:     item.BackdropPath,
				Popularity:       item.Popularity,
			}
			searchResults = append(searchResults, sr)
		}

		filtered := utils.ApplyFilters(searchResults, filters)

		// Reconstruir ResultsAll
		all.Results = make([]types.ResultsAll, len(filtered))
		for i, v := range filtered {
			all.Results[i] = types.ResultsAll{
				Id:               v.ID,
				Title:            v.Title,
				Name:             v.Name,
				MediaType:        v.MediaType,
				Overview:         v.Overview,
				VoteAverage:      v.VoteAverage,
				OriginalLanguage: v.OriginalLanguage,
				GenreIds:         v.GenreIDs,
				PosterPath:       v.PosterPath,
				BackdropPath:     v.BackdropPath,
				Popularity:       v.Popularity,
			}
		}
		all.TotalResults = len(filtered)
	}

	return all, nil
}

func FetchSearchCelebrities(language, query string, include_adult bool, page string) (types.SearchCelebrities, error) {
	var celebrities types.SearchCelebrities

	url := fmt.Sprintf("/search/person?language=%s&query=%s&include_adult=%t&page=%s", language, query, include_adult, page)

	statusCode, err := utils.FetchFromTMDB(url, &celebrities)
	if err != nil {
		return types.SearchCelebrities{}, utils.HandleTMDBError(statusCode, "search celebrities for query "+query, err)
	}

	return celebrities, nil
}

func FetchSearchCompanies(query, page string) (types.SearchCompanies, error) {
	var companies types.SearchCompanies

	url := fmt.Sprintf("/search/company?query=%s&page=%s", query, page)

	statusCode, err := utils.FetchFromTMDB(url, &companies)
	if err != nil {
		return types.SearchCompanies{}, utils.HandleTMDBError(statusCode, "search companies for query "+query, err)
	}

	return companies, nil
}
