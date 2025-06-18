package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ValidateQueryLanguage(c *gin.Context) string {
	language := c.Query("language")
	if language == "" {
		language = "en-US" // default language
	}
	return language
}

func ValidateQueryPage(c *gin.Context) string {
	page := c.Query("page")
	if page == "" {
		page = "1" // default page
	}
	return page
}

func ValidateIncludeAdult(c *gin.Context) bool {
	include_adult := c.Query("include_adult")
	if include_adult == "" {
		include_adult = "false"
	}
	return include_adult == "true"
}

func ValidateTimeWindow(timeWindow string) (string, error) {
	if timeWindow == "" {
		return "day", nil
	}

	if timeWindow != "day" && timeWindow != "week" {
		return "", fmt.Errorf("invalid time window")
	}

	return timeWindow, nil
}
