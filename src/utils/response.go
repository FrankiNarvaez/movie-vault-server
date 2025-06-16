package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleResponseOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func HandleResponse(c *gin.Context, code int, data any) {
	c.JSON(code, gin.H{
		"success": true,
		"data":    data,
	})
}
