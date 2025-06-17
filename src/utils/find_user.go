package utils

import (
	"fmt"
	config "movie/internal"
	"movie/src/models"
)

// Find user by id, username or email
func FindUser(login string) (models.User, error) {
	var user models.User
	db := config.DB

	if login == "" {
		return models.User{}, fmt.Errorf("username or email could not be empty")
	}

	query := `SELECT * FROM users WHERE id=$1 OR username=$1 OR email=$1`

	if err := db.Get(&user, query, login); err != nil {
		return models.User{}, nil
	}

	return user, nil
}
