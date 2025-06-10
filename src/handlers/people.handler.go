package handlers

import (
	"movie/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func validateQueryLanguage(c *gin.Context) string {
	language := c.Query("language")
	if language == "" {
		language = "en-US" // default language
	}
	return language
}

func GetPopularPeople(c *gin.Context) {
	language := validateQueryLanguage(c)

	people, err := services.FetchPopularPeople(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, people)
}
func GetTrendingPeople(c *gin.Context) {
	language := validateQueryLanguage(c)
	people, err := services.FetchTrendingPeople(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch people"})
		return
	}

	c.JSON(http.StatusOK, people)
}

func GetPersonDetails(c *gin.Context) {
	person_id := c.Param("person_id")
	person, err := services.FetchPersonDetails(person_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person details"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func GetPersonImages(c *gin.Context) {
	person_id := c.Param("person_id")
	images, err := services.FetchPersonImages(person_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person images"})
		return
	}

	c.JSON(http.StatusOK, images)
}

func GetPersonMovieCredits(c *gin.Context) {
	person_id := c.Param("person_id")
	credits, err := services.FetchPersonMovieCredits(person_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person movie credits"})
		return
	}

	c.JSON(http.StatusOK, credits)
}

func GetPersonCombinedCredits(c *gin.Context) {
	person_id := c.Param("person_id")
	credits, err := services.FetchPersonCombinedCredits(person_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person combined credits"})
		return
	}

	c.JSON(http.StatusOK, credits)
}

func GetPersonExternalIds(c *gin.Context) {
	person_id := c.Param("person_id")
	external_ids, err := services.FetchPersonExternalIds(person_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person external ids"})
		return
	}

	c.JSON(http.StatusOK, external_ids)
}
