package models

type Person struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          string   `json:"biography"`
	Birthday           string   `json:"birthday"`
	Deatday            *string  `json:"deathday"`
	Gender             int32    `json:"gender"`
	Homepage           *string  `json:"homepage"`
	ID                 uint32   `json:"id"`
	ImdbId             string   `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       string   `json:"place_of_birth"`
	Popularity         float32  `json:"popularity"`
	ProfilePath        string   `json:"profile_path"`
}
