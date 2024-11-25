// register_handler.go

package handler

import (
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/services"
	"strconv"

	"sinno-server/pkg/utils/typing"

	"github.com/gin-gonic/gin"
)

// GetRegistrationStatus handles retrieving registration status
func GetRegistrationStatus(c *gin.Context, queries *db.Queries) {
	// Retrieve the session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session", "details": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get member ID", "details": err.Error()})
		return
	}

	// Get activity ID from URL
	activityIDStr := c.Param("id")
	activityID, err := strconv.ParseInt(activityIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID", "details": err.Error()})
		return
	}

	// Check if the user is registered for this activity
	isRegistered, err := services.GetRegistrationStatusService(queries, int32(activityID), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check registration status", "details": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session", "details": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get member ID", "details": err.Error()})
		return
	}

	// Get activity ID from URL
	activityIDStr := c.Param("id")
	activityID, err := strconv.ParseInt(activityIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID", "details": err.Error()})
		return
	}

	// Bind JSON to CreateRegistrationParams
	var registrationData db.InsertRegistrationParams
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit registration", "details": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "Registration submitted successfully"})
}

// GetActivityRegistration handles retrieving all registrations for an activity
func GetActivityRegistration(c *gin.Context, queries *db.Queries) {
	// Get activity ID from URL
	activityIDStr := c.Param("id")
	activityID, err := strconv.ParseInt(activityIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID", "details": err.Error()})
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

func GetSubmittedMembers(c *gin.Context, queries *db.Queries) {
	// Retrieve the session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
		return
	}

	// Get user information and role from session
	userInfo, userOk := session.Values["user"].(UserInfo)
	_, roleOk := session.Values["role"].(string)
	if !userOk || !roleOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	email := userInfo.Email

	// Get member ID from the service
	memberID, err := services.GetMemberIDByEmailService(queries, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get member ID"})
		return
	}

	// Get activity ID from URL params
	activityIDStr := c.Param("id")
	activityID, err := strconv.Atoi(activityIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	// Check if the member is the proposer of the activity
	isProposer, err := services.CheckProposerService(queries, int32(activityID), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify proposer status"})
		return
	}
	if !isProposer {
		_, err := services.GetAdminIDByEmailService(queries, email)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to view this activity's registration"})
			return
		}
	}

	// Fetch registered members for the activity
	members, err := services.GetSubmittedMembersService(queries, int32(activityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch registered members"})
		return
	}

	// Return the registered members
	c.JSON(http.StatusOK, members)
}

func GetEngagements(c *gin.Context, queries *db.Queries) {
	// Retrieve the session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
		return
	}

	// Get user information and role from session
	userInfo, userOk := session.Values["user"].(UserInfo)
	_, roleOk := session.Values["role"].(string)
	if !userOk || !roleOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	email := userInfo.Email

	// Get member ID from the service
	memberID, err := services.GetMemberIDByEmailService(queries, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get member ID"})
		return
	}

	// Fetch activities for the member
	engagements, err := services.GetEngagements(queries, memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch member activities"})
		return
	}

	// Convert each engagement to the correct structure using ConvertToEngagement
	engagementList := []typing.Engagement{}
	for _, query := range engagements {
		engagement, err := typing.ConvertToEngagement(db.ListEngagementsRow(query))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to format engagement"})
			return
		}
		engagementList = append(engagementList, engagement)
	}

	// Return the member's engagements
	c.JSON(http.StatusOK, engagementList)
}
