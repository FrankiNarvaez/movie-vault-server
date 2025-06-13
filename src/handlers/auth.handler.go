package handlers

import (
	"fmt"
	config "movie/internal"
	"movie/src/errors"
	"movie/src/models"
	"movie/src/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(username string) (string, error) {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1200).Unix(),
	})

	return generateToken.SignedString([]byte(config.Secret))
}

func ValidateToken(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Secret), nil
	})

	return err == nil
}

func Login(c *gin.Context) {
	var credentials Credential
	var user models.User

	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.HandleError(c, errors.NewError("UNPROCESSABLE_ENTITY", "Invalid json", http.StatusUnprocessableEntity))
		return
	}

	if credentials.Login == "" {
		utils.HandleError(c, errors.NewBadRequestError("Username or email is required"))
		return
	}

	var err error

	user, err = utils.FindUser(credentials.Login)
	if err != nil {
		utils.HandleError(c, errors.NewError("UNAUTHORIZED", "Invalid credentials", http.StatusUnauthorized))
		return
	}

	// Verify password
	if !CheckPasswordHash(credentials.Password, user.Password) {
		utils.HandleError(c, errors.NewError("UNAUTHORIZED", "Invalid credentials", http.StatusUnauthorized))
		return
	}

	// Generate token
	token, err := GenerateToken(user.Username)
	if err != nil {
		utils.HandleError(c, errors.NewError("UNAUTHORIZED", "Could not create token", http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"token": token,
	})
}

func Register(c *gin.Context) {
	var credentials RegisterInput
	db := config.DB

	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.HandleError(c, errors.NewError("UNPROCESSABLE_ENTITY", "Invalid json", http.StatusUnprocessableEntity))
		return
	}

	if _, err := utils.FindUser(credentials.Username); err != nil {
		utils.HandleError(c, errors.NewError("CONFLICT", "User already exists", http.StatusConflict))
		return
	}

	passwordHash, err := HashPassword(credentials.Password)
	if err != nil {
		utils.HandleError(c, errors.NewBadRequestError("Failed to hash password"))
		return
	}

	queryCreateUser := `INSERT INTO users (username, email, password, created_at) values ($1, $2, $3, $4)`

	if _, err := db.Query(queryCreateUser, credentials.Username, credentials.Email, passwordHash, time.Now()); err != nil {
		utils.HandleError(c, errors.NewInternalServerError("Could not create user"))
		return
	}
	//
	// Generate token
	token, err := GenerateToken(credentials.Username)
	if err != nil {
		utils.HandleError(c, errors.NewError("UNAUTHORIZED", "Could not create token", http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": gin.H{
			"username": credentials.Username,
			"email":    credentials.Email,
		},
		"token": token,
	})
}
