package handler

import (
	"fmt"
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/services" // Import the services package

	"github.com/gin-gonic/gin"
)

// FetchAdminID handles the retrieval of an admin ID based on the provided email.
func FetchAdminID(c *gin.Context, queries *db.Queries) {
	// Retrieve session
	session, err := sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to retrieve session"})
		return
	}

	// Get user info from session
	userInfo, ok := session.Values["user"].(UserInfo)
	if !ok || userInfo.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in or session invalid"})
		return
	}

	// Use the service to fetch the admin ID
	adminID, err := services.FetchAdminIDService(queries, userInfo.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch admin ID: %s", err.Error())})
		return
	}

	// Store adminID in session
	session.Values["adminID"] = adminID
	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin ID fetched successfully", "adminID": adminID})
}
