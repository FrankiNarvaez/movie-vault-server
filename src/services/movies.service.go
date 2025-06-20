package services

import (
	"fmt"
	"movie/src/models"
	"movie/src/types"
	"movie/src/utils"
)

func FetchPopularMovies(language, page string) (types.Movies, error) {
	var movies types.Movies

	if language == "" {
		language = "en-US"
	}

	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("/movie/popular?language=%s&page=%s", language, page)
	statusCode, err := utils.FetchFromTMDB(url, &movies)
	if err != nil {
		return types.Movies{}, utils.HandleTMDBError(statusCode, "popular movies", err)
	}

	return movies, nil
}

func FetchMovieDetails(movie_id, language string) (models.Movie, error) {
	var movie models.Movie

	url := fmt.Sprintf("/movie/%s?language=%s", movie_id, language)

	statusCode, err := utils.FetchFromTMDB(url, &movie)
	if err != nil {
		return models.Movie{}, utils.HandleTMDBError(statusCode, "movie details for ID "+movie_id, err)
	}

	return movie, nil
}

func FetchMovieCredits(movie_id, language string) (types.MovieCredits, error) {
	var credits types.MovieCredits

	if language == "" {
		language = "en-US"
	}

	url := fmt.Sprintf("/movie/%s/credits?language=%s", movie_id, language)

	statusCode, err := utils.FetchFromTMDB(url, &credits)
	if err != nil {
		return types.MovieCredits{}, utils.HandleTMDBError(statusCode, "movie credits for ID "+movie_id, err)
	}

	return credits, nil
}

func FetchMovieImages(movie_id string) (types.MediaImages, error) {
	var images types.MediaImages

	url := fmt.Sprintf("/movie/%s/images", movie_id)

	statusCode, err := utils.FetchFromTMDB(url, &images)
	if err != nil {
		return types.MediaImages{}, utils.HandleTMDBError(statusCode, "movie images for ID "+movie_id, err)
	}

	return images, nil
}

func FetchMovieVideos(movie_id, language string) (types.VideoResults, error) {
	var videos types.VideoResults

	if language == "" {
		language = "en-US"
	}

	url := fmt.Sprintf("/movie/%s/videos?language=%s", movie_id, language)

	statusCode, err := utils.FetchFromTMDB(url, &videos)
	if err != nil {
		return types.VideoResults{}, utils.HandleTMDBError(statusCode, "movie videos for ID "+movie_id, err)
	}

	return videos, nil
}

func FetchMovieRecommendations(movie_id, language, page string) (types.Movies, error) {
	var recommendations types.Movies

	if language == "" {
		language = "en-US"
	}

	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("/movie/%s/recommendations?language=%s&page=%s", movie_id, language, page)

	statusCode, err := utils.FetchFromTMDB(url, &recommendations)
	if err != nil {
		return recommendations, utils.HandleTMDBError(statusCode, "movie recommendations for ID "+movie_id, err)
	}

	return recommendations, nil
}

func FetchMovieExternalIds(movie_id string) (types.ExternalIDs, error) {
	var external_ids types.ExternalIDs

	url := fmt.Sprintf("/movie/%s/external_ids", movie_id)

	statusCode, err := utils.FetchFromTMDB(url, &external_ids)
	if err != nil {
		return types.ExternalIDs{}, utils.HandleTMDBError(statusCode, "movie external ids for ID "+movie_id, err)
	}

	return external_ids, nil
}

func FetchMovieWatchProviders(movie_id string) (types.WatchProviders, error) {
	var providers types.WatchProviders

	url := fmt.Sprintf("/movie/%s/watch/providers", movie_id)

	statusCode, err := utils.FetchFromTMDB(url, &providers)
	if err != nil {
		return types.WatchProviders{}, utils.HandleTMDBError(statusCode, "movie watch providers for ID "+movie_id, err)
	}

	return providers, nil
}
