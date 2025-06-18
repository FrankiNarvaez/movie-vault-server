package services

import (
	"fmt"
	"movie/src/types"
	"movie/src/utils"
)

func FetchAllTrendings(timeWindow, language string) (types.TrendingResults, error) {
	var results types.TrendingResults

	url := fmt.Sprintf("/trending/all/%s?language=%s", timeWindow, language)
	statusCode, err := utils.FetchFromTMDB(url, &results)
	if err != nil {
		return types.TrendingResults{}, utils.HandleTMDBError(statusCode, "trending all", err)
	}

	return results, nil
}

func FetchTrendingMovies(timeWindow, language string) (types.TrendingResults, error) {
	var movies types.TrendingResults

	url := fmt.Sprintf("/trending/movie/%s?language=%s", timeWindow, language)

	statusCode, err := utils.FetchFromTMDB(url, &movies)
	if err != nil {
		return types.TrendingResults{}, utils.HandleTMDBError(statusCode, "trending movies", err)
	}

	return movies, nil
}

func FetchTrendingCelebrities(timeWindow, language string) (types.TrendingCelebrities, error) {
	var celebrities types.TrendingCelebrities
	url := fmt.Sprintf("/trending/person/%s?language=%s", timeWindow, language)

	status, err := utils.FetchFromTMDB(url, &celebrities)
	if err != nil {
		return types.TrendingCelebrities{}, utils.HandleTMDBError(status, "trending celebrities", err)
	}

	return celebrities, nil
}

func FetchTrendingSeries(timeWindow, language string) (types.TrendingResults, error) {
	var series types.TrendingResults
	url := fmt.Sprintf("/trending/tv/%s?language=%s", timeWindow, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.TrendingResults{}, utils.HandleTMDBError(status, "trending series", err)
	}
	return series, nil
}
