package utils

import (
	"encoding/json"
	"fmt"
	"io"
	config "movie/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

var baseUrl = "https://api.themoviedb.org/3"

func FetchFromTMDB(endpoint string, results any) (int, error) {
	// Create request
	req, err := http.NewRequest("GET", baseUrl+endpoint, nil)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", config.TMDBApiKey)

	// Make HTTP petition
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer res.Body.Close()

	// Verify status code
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, fmt.Errorf("API returned status code: %d", res.StatusCode)
	}

	// Read response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error reading response body: %w", err)
	}

	// Deserialize JSON from a Movie structure
	err = json.Unmarshal(body, results)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return http.StatusOK, nil
}

func ValidateQueryLanguage(c *gin.Context) string {
	language := c.Query("language")
	if language == "" {
		language = "en-US" // default language
	}
	return language
}

func ValidateQueryPage(c *gin.Context) string {
	page := c.Query("page")
	if page == "" {
		page = "1" // default page
	}
	return page
}

func ValidateIncludeAdult(c *gin.Context) bool {
	include_adult := c.Query("include_adult")
	if include_adult == "" {
		include_adult = "false"
	}
	return include_adult == "true"
}
