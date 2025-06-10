package services

import (
	"fmt"
	"movie/src/models"
	"movie/src/utils"
)

type Movies struct {
	Page         int64          `json:"page"`
	Results      []models.Movie `json:"results"`
	TotalPages   int64          `json:"total_pages"`
	TotalResults int64          `json:"total_results"`
}

type VideoResults struct {
	ID      int            `json:"id"`
	Results []models.Video `json:"results"`
}

func FetchPopularMovies(language, page string) (Movies, error) {
	var movies Movies

	if language == "" {
		language = "en-US"
	}

	if page == "" {
		page = "1"
	}

	url := fmt.Sprintf("/movie/popular?language=%s&page=%s", language, page)
	if err := utils.FetchFromTMDB(url, &movies); err != nil {
		return movies, err
	}

	return movies, nil
}

func FetchTrendingMovies(time_window, language string) (Movies, error) {
	var movies Movies

	if language == "" {
		language = "en-US"
	}

	if time_window == "" {
		time_window = "day"
	}

	url := fmt.Sprintf("/trending/movie/%s?language=%s", time_window, language)

	if err := utils.FetchFromTMDB(url, &movies); err != nil {
		return movies, err
	}

	return movies, nil
}

func FetchMovieDetails(movie_id, language string) (models.Movie, error) {
	var movie models.Movie

	url := fmt.Sprintf("/movie/%s?language=%s", movie_id, language)

	if err := utils.FetchFromTMDB(url, &movie); err != nil {
		return movie, err
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
		return credits, err
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
		return images, err
	}

	return images, nil
}

func FetchMovieVideos(movie_id, language string) (VideoResults, error) {
	var videos VideoResults

	if language == "" {
		language = "en-US"
	}

	url := fmt.Sprintf("/movie/%s/videos?language=%s", movie_id, language)

	if err := utils.FetchFromTMDB(url, &videos); err != nil {
		return videos, err
	}

	return videos, nil
}

func FetchMovieRecommendations(movie_id, language, page string) (Movies, error) {
	var recommendations Movies

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

func FetchMovieExternalIds(movie_id string) (models.ExternalIDs, error) {
	var external_ids models.ExternalIDs

	url := fmt.Sprintf("/movie/%s/external_ids", movie_id)

	if err := utils.FetchFromTMDB(url, &external_ids); err != nil {
		return external_ids, err
	}

	return external_ids, nil
}

func FetchMovieWatchProviders(movie_id string) (models.WatchProviders, error) {
	var providers models.WatchProviders

	url := fmt.Sprintf("/movie/%s/external_ids", movie_id)

	if err := utils.FetchFromTMDB(url, &providers); err != nil {
		return providers, err
	}

	return providers, nil
}
