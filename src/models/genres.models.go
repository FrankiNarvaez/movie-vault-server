package models

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Genres struct {
	Genres []Genre `json:"genres"`
}
