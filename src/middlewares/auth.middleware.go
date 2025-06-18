package middlewares

import (
	"movie/src/errors"
	"movie/src/handlers"
	"movie/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		utils.HandleError(c, errors.NewError("UNAUTHORIZED", "Authorization header is missing", http.StatusUnauthorized))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		utils.HandleError(c, errors.NewError("UNAUTHORIZED", "Invalid token format", http.StatusUnauthorized))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := handlers.ValidateToken(authToken[1])
	if err != nil {
		utils.HandleError(c, errors.NewError("UNAUTHORIZED", err.Error(), http.StatusUnauthorized))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	c.Next()
}
