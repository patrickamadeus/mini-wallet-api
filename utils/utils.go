package utils

import (
	"time"
	"github.com/gin-gonic/gin"
)

func CurrentTime() time.Time {
	return time.Now().UTC()
}

func ErrorResponse(c *gin.Context, statusCode int, errors map[string][]string) {
	c.JSON(statusCode, gin.H{
		"status": "fail",
		"data": gin.H{
			"error": errors,
		},
	})
}

func SuccessResponse(c *gin.Context, statusCode int, data gin.H) {
	c.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
	})
}