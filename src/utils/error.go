package utils

import (
	"fmt"
	"movie/src/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if customErr, ok := err.(errors.CustomError); ok {
		c.JSON(customErr.StatusCode, gin.H{
			"type":        customErr.Type,
			"error":       customErr.Message,
			"status_code": customErr.StatusCode,
		})
		return
	}

	// Generic error
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error",
		"code":  http.StatusInternalServerError,
	})
}

func HandleTMDBError(statusCode int, resource string, originalErr error) error {
	switch statusCode {
	case http.StatusNotFound:
		return errors.NewNotFoundError(fmt.Sprintf("Could not find %s", resource))
	case http.StatusBadRequest:
		return errors.NewBadRequestError(fmt.Sprintf("Invalid request for %s", resource))
	case http.StatusTooManyRequests:
		return errors.NewExternalAPIError("API rate limit exceeded")
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable:
		return errors.NewExternalAPIError(fmt.Sprintf("TMDB service temporarily unavailable for %s", resource))
	default:
		return errors.NewInternalServerError(fmt.Sprintf("Failed to fetch %s: %v", resource, originalErr))
	}
}
