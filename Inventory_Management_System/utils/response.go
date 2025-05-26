package utils

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// JSONErrorResponse sends a JSON error response with specified status code and message
func JSONErrorResponse(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, gin.H{"error": message})
    c.Abort()
}

// JSONSuccessResponse sends a JSON success response with data
func JSONSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
    c.JSON(statusCode, data)
}
