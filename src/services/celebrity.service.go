package services

import (
	"fmt"
	"movie/src/models"
	"movie/src/types"
	"movie/src/utils"
)

func FetchPopularPeople(language string, page string) (types.PopularPeople, error) {
	var people types.PopularPeople
	url := fmt.Sprintf("/person/popular?language=%s&page=%s", language, page)

	status, err := utils.FetchFromTMDB(url, &people)
	if err != nil {
		return types.PopularPeople{}, utils.HandleTMDBError(status, "popular people", err)
	}

	return people, nil
}

func FetchTrendingPeople(language string, time_window string) (types.TrendingPeople, error) {
	var people types.TrendingPeople
	url := fmt.Sprintf("/trending/person/%s?language=%s", time_window, language)

	status, err := utils.FetchFromTMDB(url, &people)
	if err != nil {
		return types.TrendingPeople{}, utils.HandleTMDBError(status, "trending people", err)
	}

	return people, nil
}

func FetchPersonDetails(person_id string, language string) (models.Person, error) {
	var person models.Person
	url := fmt.Sprintf("/person/%s?language=%s", person_id, language)

	status, err := utils.FetchFromTMDB(url, &person)
	if err != nil {
		return models.Person{}, utils.HandleTMDBError(status, "person details for ID "+person_id, err)
	}

	return person, nil
}

func FetchPersonImages(person_id string) (types.PersonImages, error) {
	var images types.PersonImages
	url := fmt.Sprintf("/person/%s/images", person_id)

	status, err := utils.FetchFromTMDB(url, &images)
	if err != nil {
		return types.PersonImages{}, utils.HandleTMDBError(status, "person images for ID "+person_id, err)
	}

	return images, nil
}

func FetchPersonMovieCredits(person_id string, language string) (types.PersonMovieCredits, error) {
	var credits types.PersonMovieCredits
	url := fmt.Sprintf("/person/%s/movie_credits?language=%s", person_id, language)

	status, err := utils.FetchFromTMDB(url, &credits)
	if err != nil {
		return types.PersonMovieCredits{}, utils.HandleTMDBError(status, "person movie credits for ID "+person_id, err)
	}

	return credits, nil
}

func FetchPersonCombinedCredits(person_id string, language string) (types.PersonCombinedCredits, error) {
	var credits types.PersonCombinedCredits
	url := fmt.Sprintf("/person/%s/combined_credits?language=%s", person_id, language)

	status, err := utils.FetchFromTMDB(url, &credits)
	if err != nil {
		return types.PersonCombinedCredits{}, utils.HandleTMDBError(status, "person combined credits for ID "+person_id, err)
	}

	return credits, nil
}

func FetchPersonExternalIds(person_id string) (types.ExternalIDs, error) {
	var ids types.ExternalIDs
	url := fmt.Sprintf("/person/%s/external_ids", person_id)

	status, err := utils.FetchFromTMDB(url, &ids)
	if err != nil {
		return types.ExternalIDs{}, utils.HandleTMDBError(status, "person external ids for ID "+person_id, err)
	}

	return ids, nil
}
