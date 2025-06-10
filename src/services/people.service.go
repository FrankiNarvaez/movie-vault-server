package services

import (
	"encoding/json"
	"fmt"
	"io"
	config "movie/internal"
	"movie/src/models"
	"movie/src/types"
	"net/http"
)

var url_base string = "https://api.themoviedb.org/3"

func fetchFromTMDB(final_url string, result interface{}) error {
	req, err := http.NewRequest("GET", url_base+final_url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", config.TMDBApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	return nil
}

func FetchPopularPeople(language string, page string) (types.PopularPeople, error) {
	var people types.PopularPeople
	url := fmt.Sprintf("/person/popular?language=%s&page=%s", language, page)

	err := fetchFromTMDB(url, &people)
	if err != nil {
		return types.PopularPeople{}, err
	}

	return people, nil
}

func FetchTrendingPeople(language string, time_window string) (types.TrendingPeople, error) {
	var people types.TrendingPeople
	url := fmt.Sprintf("/trending/person/%s?language=%s", time_window, language)

	err := fetchFromTMDB(url, &people)
	if err != nil {
		return types.TrendingPeople{}, err
	}

	return people, nil
}

func FetchPersonDetails(person_id string, language string) (models.Person, error) {
	var person models.Person
	url := fmt.Sprintf("/person/%s?language=%s", person_id, language)

	err := fetchFromTMDB(url, &person)
	if err != nil {
		return models.Person{}, err
	}

	return person, nil
}

func FetchPersonImages(person_id string) (types.PersonImages, error) {
	var images types.PersonImages
	url := fmt.Sprintf("/person/%s/images", person_id)

	err := fetchFromTMDB(url, &images)
	if err != nil {
		return types.PersonImages{}, err
	}

	return images, nil
}

func FetchPersonMovieCredits(person_id string, language string) (types.PersonMovieCredits, error) {
	var credits types.PersonMovieCredits
	url := fmt.Sprintf("/person/%s/movie_credits?language=%s", person_id, language)

	err := fetchFromTMDB(url, &credits)
	if err != nil {
		return types.PersonMovieCredits{}, err
	}

	return credits, nil
}

func FetchPersonCombinedCredits(person_id string, language string) (types.PersonCombinedCredits, error) {
	var credits types.PersonCombinedCredits
	url := fmt.Sprintf("/person/%s/combined_credits?language=%s", person_id, language)

	err := fetchFromTMDB(url, &credits)
	if err != nil {
		return types.PersonCombinedCredits{}, err
	}

	return credits, nil
}

func FetchPersonExternalIds(person_id string) (types.ExternalIDs, error) {
	var ids types.ExternalIDs
	url := fmt.Sprintf("/person/%s/external_ids", person_id)

	err := fetchFromTMDB(url, &ids)
	if err != nil {
		return types.ExternalIDs{}, err
	}

	return ids, nil
}
