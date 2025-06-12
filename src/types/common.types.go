package types

type ExternalIDs struct {
	Id           int     `json:"id"`
	Freebase_mid *string `json:"freebase_mid"`
	Freebase_sid *string `json:"freebase_sid"`
	Freebase_id  *string `json:"freebase_id"`
	Imdb_id      *string `json:"imdb_id"`
	Tvrage_id    *int64  `json:"tvrage_id"`
	TvdbId       *int64  `json:"tvdb_id"`
	Wikidata_id  *string `json:"wikidata_id"`
	Facebook_id  *string `json:"facebook_id"`
	Instagram_id *string `json:"instagram_id"`
	Tiktok_id    *string `json:"tiktok_id"`
	Twitter_id   *string `json:"twitter_id"`
	Youtube_id   *string `json:"youtube_id"`
}
