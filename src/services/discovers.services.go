package services

import (
	"fmt"
	"movie/src/types"
	"movie/src/utils"
)

func FetchDiscoverMovies(includeAdult, includeVideo bool, language, page, sortBy, withCompanies, withGenres string) (*types.DiscoverMoviesResults, error) {
	var movies types.DiscoverMoviesResults

	url := fmt.Sprintf("/discover/movie?include_adult=%t&include_video=%t&language=%s&page=%s&sort_by=%s&with_companies=%s&with_genres=%s", includeAdult, includeVideo, language, page, sortBy, withCompanies, withGenres)
	statusCode, err := utils.FetchFromTMDB(url, &movies)
	if err != nil {
		return nil, utils.HandleTMDBError(statusCode, "discover movies", err)
	}
	return &movies, nil
}

func FetchDiscoverSeries(includeAdult, includeNullFirstAirDates bool, language, page, sortBy, withCompanies, withGenres, withNetworks string) (*types.DiscoverSeriesResults, error) {
	var series types.DiscoverSeriesResults

	url := fmt.Sprintf("/discover/tv?include_adult=%t&include_video=%t&language=%s&page=%s&sort_by=%s&with_companies=%s&with_genres=%s&with_networks=%s", includeAdult, includeNullFirstAirDates, language, page, sortBy, withCompanies, withGenres, withNetworks)
	statusCode, err := utils.FetchFromTMDB(url, &series)
	if err != nil {
		return nil, utils.HandleTMDBError(statusCode, "discover series", err)
	}
	return &series, nil
}
