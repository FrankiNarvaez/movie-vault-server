package services

import (
	"fmt"
	"movie/src/types"
	"movie/src/utils"
)

func FetchAllTrendings(time_window, language string) (types.TrendingResults, error) {
	var results types.TrendingResults

	return results, nil
}

func FetchTrendingMovies(time_window, language string) (types.Movies, error) {
	var movies types.Movies

	if language == "" {
		language = "en-US"
	}

	url := fmt.Sprintf("/trending/movie/%s?language=%s", time_window, language)

	statusCode, err := utils.FetchFromTMDB(url, &movies)
	if err != nil {
		return types.Movies{}, utils.HandleTMDBError(statusCode, "trending movies", err)
	}

	return movies, nil
}

func FetchTrendingCelebrities(language string, time_window string) (types.TrendingCelebrities, error) {
	var celebrities types.TrendingCelebrities
	url := fmt.Sprintf("/trending/person/%s?language=%s", time_window, language)

	status, err := utils.FetchFromTMDB(url, &celebrities)
	if err != nil {
		return types.TrendingCelebrities{}, utils.HandleTMDBError(status, "trending celebrities", err)
	}

	return celebrities, nil
}

func FetchTrendingSeries(language, time string) (types.TrendingResults, error) {
	var series types.TrendingResults
	url := fmt.Sprintf("/trending/tv/%s?language=%s", time, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.TrendingResults{}, utils.HandleTMDBError(status, "trending series", err)
	}
	return series, nil
}
