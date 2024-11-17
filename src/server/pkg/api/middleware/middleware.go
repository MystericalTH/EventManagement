package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Extract token from headers
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Validate token (pseudo-code)
	if !isValidToken(token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next() // Continue to the next handler
}

// Example of token validation function
func isValidToken(token string) bool {
	// Implement your token validation logic here
	return token == "your_secret_token" // Example
}
