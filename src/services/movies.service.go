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
	if err := utils.FetchFromTMDB(url, &movies); err != nil {
		return types.Movies{}, err
	}

	return movies, nil
}

func FetchTrendingMovies(time_window, language string) (types.Movies, error) {
	var movies types.Movies

	if language == "" {
		language = "en-US"
	}

	if time_window == "" {
		time_window = "day"
	}

	url := fmt.Sprintf("/trending/movie/%s?language=%s", time_window, language)

	if err := utils.FetchFromTMDB(url, &movies); err != nil {
		return types.Movies{}, err
	}

	return movies, nil
}

func FetchMovieDetails(movie_id, language string) (models.Movie, error) {
	var movie models.Movie

	url := fmt.Sprintf("/movie/%s?language=%s", movie_id, language)

	if err := utils.FetchFromTMDB(url, &movie); err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

func FetchMovieCredits(movie_id, language string) (models.Movie, error) {
	var credits models.Movie

	if language == "" {
		language = "en-US"
	}

	url := fmt.Sprintf("/movie/%s/credits?language=%s", movie_id, language)

	if err := utils.FetchFromTMDB(url, &credits); err != nil {
		return models.Movie{}, err
	}

	return credits, nil
}

func FetchMovieImages(movie_id, language string) (models.MediaImages, error) {
	var images models.MediaImages

	if language == "" {
		language = "en-US"
	}

	url := fmt.Sprintf("/movie/%s/images?language=%s", movie_id, language)

	if err := utils.FetchFromTMDB(url, &images); err != nil {
		return models.MediaImages{}, err
	}

	return images, nil
}

func FetchMovieVideos(movie_id, language string) (types.VideoResults, error) {
	var videos types.VideoResults

	if language == "" {
		language = "en-US"
	}

	url := fmt.Sprintf("/movie/%s/videos?language=%s", movie_id, language)

	if err := utils.FetchFromTMDB(url, &videos); err != nil {
		return types.VideoResults{}, err
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

	if err := utils.FetchFromTMDB(url, &recommendations); err != nil {
		return recommendations, err
	}

	return recommendations, nil
}

func FetchMovieExternalIds(movie_id string) (types.ExternalIDs, error) {
	var external_ids types.ExternalIDs

	url := fmt.Sprintf("/movie/%s/external_ids", movie_id)

	if err := utils.FetchFromTMDB(url, &external_ids); err != nil {
		return types.ExternalIDs{}, err
	}

	return external_ids, nil
}

func FetchMovieWatchProviders(movie_id string) (models.WatchProviders, error) {
	var providers models.WatchProviders

	url := fmt.Sprintf("/movie/%s/external_ids", movie_id)

	if err := utils.FetchFromTMDB(url, &providers); err != nil {
		return models.WatchProviders{}, err
	}

	return providers, nil
}
