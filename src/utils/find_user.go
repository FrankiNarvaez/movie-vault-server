package utils

import (
	"fmt"
	config "movie/internal"
	"movie/src/models"

	"github.com/gin-gonic/gin"
)

// Find user by username or email
func FindUser(login string) (models.User, error) {
	var user models.User
	db := config.DB

	if login == "" {
		return models.User{}, fmt.Errorf("username or email could not be empty")
	}

	query := `SELECT * FROM users WHERE username=$1 OR email=$1`

	if err := db.Get(&user, query, login); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Find user by ID
func FindUserByID(id string) (models.User, error) {
	var user models.User
	db := config.DB

	if id == "" {
		return models.User{}, fmt.Errorf("id could not be empty")
	}

	if err := db.Get(&user, `SELECT * FROM users WHERE id=$1`, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Get user from gin.Context
func GetUserFromContext(c *gin.Context) (models.User, error) {
	userInterface, exists := c.Get("user")
	if !exists {
		return models.User{}, fmt.Errorf("user not authenticated")
	}

	user, ok := userInterface.(models.User)
	if !ok {
		return models.User{}, fmt.Errorf("invalid user data")
	}

	return user, nil
}
