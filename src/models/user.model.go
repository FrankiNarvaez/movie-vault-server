package models

import "time"

type User struct {
	ID         string     `json:"id" db:"id"`
	Username   string     `json:"username" db:"username"`
	Email      string     `json:"email" db:"email"`
	Password   string     `json:"-" db:"password"`
	RememberAt *time.Time `json:"remember_at" db:"remember_at"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
}
