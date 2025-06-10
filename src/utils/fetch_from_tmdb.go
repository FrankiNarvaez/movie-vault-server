package utils

import (
	"encoding/json"
	"fmt"
	"io"
	config "movie/internal"
	"net/http"
)

var baseUrl = "https://api.themoviedb.org/3"

func FetchFromTMDB(endpoint string, results any) error {
	// Create request
	req, err := http.NewRequest("GET", baseUrl+endpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", config.TMDBApiKey)

	// Make HTTP petition
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %w", err)
	}
	defer res.Body.Close()

	// Verify status code
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code: %d", res.StatusCode)
	}

	// Read response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// Deserialize JSON from a Movie structure
	err = json.Unmarshal(body, results)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return nil
}
