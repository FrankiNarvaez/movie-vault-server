package handlers

import (
	"fmt"
	config "movie/internal"
	"movie/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UserExists(credentials Credential) bool {
	db := config.DB
	var count int

	// Search by username if exists
	if credentials.Username != nil {
		query := `SELECT COUNT(*) FROM users WHERE username=$1`
		if err := db.Get(&count, query, *credentials.Username); err != nil {
			fmt.Printf("Database error: %v\n", err)
			return false
		}
		return count > 0
	}

	// Search by username if username not exists
	if credentials.Email != nil {
		query := `SELECT COUNT(*) FROM users WHERE email=$1`
		if err := db.Get(&count, query, *credentials.Email); err != nil {
			fmt.Printf("Database error: %v\n", err)
			return false
		}
		return count > 0
	}

	// return false if username and email not exits
	return false
}

func GenerateToken(username string) (string, error) {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
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
	db := config.DB

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar que al menos username o email estén presentes
	if credentials.Username == nil && credentials.Email == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email is required"})
		return
	}

	var err error

	// First search by username if exists
	if credentials.Username != nil && *credentials.Username != "" {
		query := `SELECT * FROM users WHERE username=$1`
		err = db.Get(&user, query, *credentials.Username)
	} else if credentials.Email != nil && *credentials.Email != "" {
		query := `SELECT * FROM users WHERE email=$1`
		err = db.Get(&user, query, *credentials.Email)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email cannot be empty"})
		return
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verify password
	if !CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token
	token, err := GenerateToken(*credentials.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"token":    token,
		},
	})
}

func Register(c *gin.Context) {
	var credentials Credential
	db := config.DB

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if UserExists(credentials) {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	passwordHash, err := HashPassword(credentials.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	query_create_user := `INSERT INTO users (username, email, password, created_at) values ($1, $2, $3, $4)`

	if _, err := db.Query(query_create_user, *credentials.Username, *credentials.Email, passwordHash, time.Now()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": gin.H{
		"username": *credentials.Username,
		"email":    *credentials.Email,
	}})
}

func Logout(c *gin.Context) {}
