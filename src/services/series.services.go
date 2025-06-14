package services

import (
	"fmt"
	"movie/src/models"
	"movie/src/types"
	"movie/src/utils"
)

func FetchPopularSeries(language, page string) (types.PopularTv, error) {
	var series types.PopularTv
	url := fmt.Sprintf("/tv/popular?language=%s&page=%s", language, page)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.PopularTv{}, utils.HandleTMDBError(status, "popular series", err)
	}
	return series, nil
}

func FetchTrendingSeries(language, time string) (types.TrendingTv, error) {
	var series types.TrendingTv
	url := fmt.Sprintf("/trending/tv/%s?language=%s", time, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.TrendingTv{}, utils.HandleTMDBError(status, "trending series", err)
	}
	return series, nil
}

func FetchSeriesDetails(series_id, language, append_to_response string) (types.SeriesDetails, error) {
	var series types.SeriesDetails
	url := fmt.Sprintf("/tv/%s?language=%s", series_id, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeriesDetails{}, utils.HandleTMDBError(status, "series details", err)
	}
	return series, nil
}

func FetchSeriesCredits(series_id, language string) (types.SeriesCredits, error) {
	var series types.SeriesCredits
	url := fmt.Sprintf("/tv/%s/credits?language=%s", series_id, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeriesCredits{}, utils.HandleTMDBError(status, "series credits", err)
	}
	return series, nil
}

func FetchSeriesImages(series_id string) (types.SeriesImages, error) {
	var series types.SeriesImages
	url := fmt.Sprintf("/tv/%s/images", series_id)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeriesImages{}, utils.HandleTMDBError(status, "series images", err)
	}
	return series, nil
}

func FetchSeriesVideos(series_id, language string) (types.SeriesVideos, error) {
	var series types.SeriesVideos
	url := fmt.Sprintf("/tv/%s/videos?language=%s", series_id, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeriesVideos{}, utils.HandleTMDBError(status, "series videos", err)
	}
	return series, nil
}

func FetchSeriesRecommendations(series_id, language, page string) (types.SeriesRecommendations, error) {
	var series types.SeriesRecommendations
	url := fmt.Sprintf("/tv/%s/recommendations?language=%s&page=%s", series_id, language, page)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeriesRecommendations{}, utils.HandleTMDBError(status, "series recommendations", err)
	}
	return series, nil
}

func FetchSeriesExternalIds(series_id string) (types.ExternalIDs, error) {
	var series types.ExternalIDs
	url := fmt.Sprintf("/tv/%s/external_ids", series_id)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.ExternalIDs{}, utils.HandleTMDBError(status, "series external ids", err)
	}
	return series, nil
}

func FetchSeriesWatchProviders(series_id string) (types.WatchProviders, error) {
	var series types.WatchProviders
	url := fmt.Sprintf("/tv/%s/watch/providers", series_id)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.WatchProviders{}, utils.HandleTMDBError(status, "series watch providers", err)
	}
	return series, nil
}

func FetchSeasonDetails(series_id, season_number, language string) (types.SeasonDetails, error) {
	var series types.SeasonDetails
	url := fmt.Sprintf("/tv/%s/season/%s?language=%s", series_id, season_number, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeasonDetails{}, utils.HandleTMDBError(status, "season details", err)
	}
	return series, nil
}

func FetchSeasonCredits(series_id, season_number, language string) (types.SeasonCredits, error) {
	var series types.SeasonCredits
	url := fmt.Sprintf("/tv/%s/season/%s/credits?language=%s", series_id, season_number, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeasonCredits{}, utils.HandleTMDBError(status, "season credits", err)
	}
	return series, nil
}

func FetchSeasonExternalIds(series_id, season_number string) (types.ExternalIDs, error) {
	var series types.ExternalIDs
	url := fmt.Sprintf("/tv/%s/season/%s/external_ids", series_id, season_number)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.ExternalIDs{}, utils.HandleTMDBError(status, "season external ids", err)
	}
	return series, nil
}

func FetchSeasonImages(series_id, season_number string) (types.SeriesImages, error) {
	var series types.SeriesImages
	url := fmt.Sprintf("/tv/%s/season/%s/images", series_id, season_number)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeriesImages{}, utils.HandleTMDBError(status, "season images", err)
	}
	return series, nil
}

func FetchSeasonVideos(series_id, season_number, language string) (types.SeriesVideos, error) {
	var series types.SeriesVideos
	url := fmt.Sprintf("/tv/%s/season/%s/videos?language=%s", series_id, season_number, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.SeriesVideos{}, utils.HandleTMDBError(status, "season videos", err)
	}
	return series, nil
}

func FetchSeasonWatchProviders(series_id, season_number, language string) (types.WatchProviders, error) {
	var series types.WatchProviders
	url := fmt.Sprintf("/tv/%s/season/%s/watch/providers?language=%s", series_id, season_number, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.WatchProviders{}, utils.HandleTMDBError(status, "season watch providers", err)
	}
	return series, nil
}

func FetchEpisodeDetails(series_id, season_number, episode_number, language string) (models.Episode, error) {
	var episode models.Episode
	url := fmt.Sprintf("/tv/%s/season/%s/episode/%s", series_id, season_number, episode_number)

	if status, err := utils.FetchFromTMDB(url, &episode); err != nil {
		return models.Episode{}, utils.HandleTMDBError(status, "episode details", err)
	}
	return episode, nil
}

func FetchEpisodeCredits(series_id, season_number, episode_number, language string) (types.EpisodeCredits, error) {
	var series types.EpisodeCredits
	url := fmt.Sprintf("/tv/%s/season/%s/episode/%s/credits?language=%s", series_id, season_number, episode_number, language)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.EpisodeCredits{}, utils.HandleTMDBError(status, "episode credits", err)
	}
	return series, nil
}

func FetchEpisodeExternalIds(series_id, season_number, episode_number string) (types.ExternalIDs, error) {
	var episode types.ExternalIDs
	url := fmt.Sprintf("/tv/%s/season/%s/episode/%s/external_ids", series_id, season_number, episode_number)

	if status, err := utils.FetchFromTMDB(url, &episode); err != nil {
		return types.ExternalIDs{}, utils.HandleTMDBError(status, "episode external ids", err)
	}
	return episode, nil
}

func FetchEpisodeImages(series_id, season_number, episode_number string) (types.EpisodeImages, error) {
	var series types.EpisodeImages
	url := fmt.Sprintf("/tv/%s/season/%s/episode/%s/images", series_id, season_number, episode_number)

	if status, err := utils.FetchFromTMDB(url, &series); err != nil {
		return types.EpisodeImages{}, utils.HandleTMDBError(status, "episode images", err)
	}
	return series, nil
}

func FetchEpisodeVideos(series_id, season_number, episode_number, language string) (types.SeriesVideos, error) {
	var episode types.SeriesVideos
	url := fmt.Sprintf("/tv/%s/season/%s/episode/%s/videos?language=%s", series_id, season_number, episode_number, language)

	if status, err := utils.FetchFromTMDB(url, &episode); err != nil {
		return types.SeriesVideos{}, utils.HandleTMDBError(status, "episode videos", err)
	}
	return episode, nil
}
