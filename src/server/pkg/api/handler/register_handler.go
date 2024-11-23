// register_handler.go

package handler

import (
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRegistrationStatus handles retrieving registration status
func GetRegistrationStatus(c *gin.Context, queries *db.Queries) {
	// Retrieve the session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
		return
	}

	// Get user information and role from session
	userInfo, userOk := session.Values["user"].(UserInfo)
	role, roleOk := session.Values["role"].(string)
	if !userOk || !roleOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	email := userInfo.Email

	// Only allow members to check registration status
	if role != "member" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only members can check registration status"})
		return
	}

	// Get member ID from the service
	memberID, err := services.GetMemberIDByEmailService(queries, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get member ID"})
		return
	}

	// Get activity ID from URL
	activityIDStr := c.Param("activityId")
	activityID, err := strconv.ParseInt(activityIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	// Check if the user is registered for this activity
	isRegistered, err := services.GetRegistrationStatusService(queries, int32(activityID), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check registration status"})
		return
	}

	// Return JSON response
	c.JSON(http.StatusOK, gin.H{"isRegistered": isRegistered})
}

// SubmitRegistration handles submitting a new registration
func SubmitRegistration(c *gin.Context, queries *db.Queries) {
	// Retrieve the session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
		return
	}

	// Get user information and role from session
	userInfo, userOk := session.Values["user"].(UserInfo)
	role, roleOk := session.Values["role"].(string)
	if !userOk || !roleOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	email := userInfo.Email

	// Only allow members to register
	if role != "member" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only members can register"})
		return
	}

	// Get member ID from the service
	memberID, err := services.GetMemberIDByEmailService(queries, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get member ID"})
		return
	}

	// Get activity ID from URL
	activityIDStr := c.Param("activityId")
	activityID, err := strconv.ParseInt(activityIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	// Bind JSON to CreateRegistrationParams
	var registrationData db.InsertRegistrationParams
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create registration parameters
	params := db.InsertRegistrationParams{
		Activityid:  int32(activityID),
		Memberid:    memberID,
		Role:        registrationData.Role,
		Expectation: registrationData.Expectation,
	}

	// Insert the registration using the service
	if err := services.CreateRegistrationService(queries, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit registration"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "Registration submitted successfully"})
}

// GetActivityRegistration handles retrieving all registrations for an activity
func GetActivityRegistration(c *gin.Context, queries *db.Queries) {
	// Get activity ID from URL
	activityIDStr := c.Param("activityId")
	activityID, err := strconv.ParseInt(activityIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	// Get all registrations for this activity
	registrations, err := services.GetActivityRegistrationService(queries, int32(activityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get registrations", "message": err.Error()})
		return
	}

	// Return JSON response
	c.JSON(http.StatusOK, registrations)
}
