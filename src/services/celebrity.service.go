package services

import (
	"fmt"
	"movie/src/models"
	"movie/src/types"
	"movie/src/utils"
)

func FetchPopularCelebrities(language string, page string) (types.PopularCelebrities, error) {
	var celebrities types.PopularCelebrities
	url := fmt.Sprintf("/person/popular?language=%s&page=%s", language, page)

	status, err := utils.FetchFromTMDB(url, &celebrities)
	if err != nil {
		return types.PopularCelebrities{}, utils.HandleTMDBError(status, "popular celebrities", err)
	}

	return celebrities, nil
}

func FetchTrendingCelebrities(language string, time_window string) (types.TrendingCelebrities, error) {
	var celebrities types.TrendingCelebrities
	url := fmt.Sprintf("/trending/person/%s?language=%s", time_window, language)

	status, err := utils.FetchFromTMDB(url, &celebrities)
	if err != nil {
		return types.TrendingCelebrities{}, utils.HandleTMDBError(status, "trending celebrities", err)
	}

	return celebrities, nil
}

func FetchCelebrityDetails(celebrity_id string, language string) (models.Celebrity, error) {
	var celebrity models.Celebrity
	url := fmt.Sprintf("/person/%s?language=%s", celebrity_id, language)

	status, err := utils.FetchFromTMDB(url, &celebrity)
	if err != nil {
		return models.Celebrity{}, utils.HandleTMDBError(status, "celebrity details for ID "+celebrity_id, err)
	}

	return celebrity, nil
}

func FetchCelebrityImages(celebrity_id string) (types.CelebrityImages, error) {
	var images types.CelebrityImages
	url := fmt.Sprintf("/person/%s/images", celebrity_id)

	status, err := utils.FetchFromTMDB(url, &images)
	if err != nil {
		return types.CelebrityImages{}, utils.HandleTMDBError(status, "celebrity images for ID "+celebrity_id, err)
	}

	return images, nil
}

func FetchCelebrityMovieCredits(celebrity_id string, language string) (types.CelebrityMovieCredits, error) {
	var credits types.CelebrityMovieCredits
	url := fmt.Sprintf("/person/%s/movie_credits?language=%s", celebrity_id, language)

	status, err := utils.FetchFromTMDB(url, &credits)
	if err != nil {
		return types.CelebrityMovieCredits{}, utils.HandleTMDBError(status, "celebrity movie credits for ID "+celebrity_id, err)
	}

	return credits, nil
}

func FetchCelebrityCombinedCredits(celebrity_id string, language string) (types.CelebrityCombinedCredits, error) {
	var credits types.CelebrityCombinedCredits
	url := fmt.Sprintf("/person/%s/combined_credits?language=%s", celebrity_id, language)

	status, err := utils.FetchFromTMDB(url, &credits)
	if err != nil {
		return types.CelebrityCombinedCredits{}, utils.HandleTMDBError(status, "celebrity combined credits for ID "+celebrity_id, err)
	}

	return credits, nil
}

func FetchCelebrityExternalIds(celebrity_id string) (types.ExternalIDs, error) {
	var ids types.ExternalIDs
	url := fmt.Sprintf("/person/%s/external_ids", celebrity_id)

	status, err := utils.FetchFromTMDB(url, &ids)
	if err != nil {
		return types.ExternalIDs{}, utils.HandleTMDBError(status, "celebrity external ids for ID "+celebrity_id, err)
	}

	return ids, nil
}
