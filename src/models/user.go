package models

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberAt string `json:"remember_at"`
}
