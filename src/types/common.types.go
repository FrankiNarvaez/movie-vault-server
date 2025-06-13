package types

type ExternalIDs struct {
	Id           int     `json:"id"`
	Freebase_mid *string `json:"freebase_mid,omitempty"`
	Freebase_sid *string `json:"freebase_sid,omitempty"`
	Freebase_id  *string `json:"freebase_id,omitempty"`
	Imdb_id      *string `json:"imdb_id,omitempty"`
	Tvrage_id    *int64  `json:"tvrage_id,omitempty"`
	TvdbId       *int64  `json:"tvdb_id,omitempty"`
	Wikidata_id  *string `json:"wikidata_id,omitempty"`
	Facebook_id  *string `json:"facebook_id,omitempty"`
	Instagram_id *string `json:"instagram_id,omitempty"`
	Tiktok_id    *string `json:"tiktok_id,omitempty"`
	Twitter_id   *string `json:"twitter_id,omitempty"`
	Youtube_id   *string `json:"youtube_id,omitempty"`
}
