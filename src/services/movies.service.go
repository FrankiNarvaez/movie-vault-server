package services

import (
	"encoding/json"
	"fmt"
	"io"
	config "movie/internal"
	"movie/src/models"
	"net/http"
)

type PopularMovies struct {
	Page         int64          `json:"page"`
	Results      []models.Movie `json:"results"`
	TotalPages   int64          `json:"total_pages"`
	TotalResults int64          `json:"total_results"`
}

func FetchPopularMovies(language string) (PopularMovies, error) {
	var movie PopularMovies

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?language=%s", language)
	api_key := fmt.Sprintf("%s", config.TMDBApiKey)

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return movie, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", api_key)

	// Make HTTP petition
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return movie, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer res.Body.Close()

	// Verify status code
	if res.StatusCode != http.StatusOK {
		return movie, fmt.Errorf("API returned status code: %d", res.StatusCode)
	}

	// Read response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return movie, fmt.Errorf("error reading response body: %w", err)
	}

	// Deserialize JSON from a Movie structure
	err = json.Unmarshal(body, &movie)
	if err != nil {
		return movie, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return movie, nil
}
