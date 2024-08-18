package handlers

import (
	"github.com/gin-gonic/gin"
)

func httpError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
