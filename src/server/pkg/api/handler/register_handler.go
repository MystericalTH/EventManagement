// register_handler.go

package handler

import (
	"log"
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
		log.Printf("Error retrieving session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
		return
	}

	// Get user information and role from session
	userInfo, userOk := session.Values["user"].(UserInfo)
	role, roleOk := session.Values["role"].(string)
	if !userOk || !roleOk {
		log.Println("Unauthorized access attempt: User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	email := userInfo.Email

	log.Printf("User email retrieved from session: %s", email)

	// Get activity ID from URL params
	activityIDStr := c.Param("id")
	activityID, err := strconv.Atoi(activityIDStr)
	if err != nil {
		log.Printf("Invalid activity ID: %s", activityIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}
	log.Printf("Activity ID retrieved: %d", activityID)

	// Get member ID from the service
	if role == "member" {
		memberID, err := services.GetMemberIDByEmailService(queries, email)
		if err != nil {
			log.Printf("Error fetching member ID for email %s: %v", email, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get member ID"})
			return
		}
		log.Printf("Member ID retrieved: %d", memberID)
		_, err = services.CheckProposerService(queries, int32(activityID), memberID)
		if err != nil {
			log.Printf("User %s is not the proposer for activity ID %d. Checking admin access.", email, activityID)

		}
	} else if role == "admin" {
		_, err := services.GetAdminIDByEmailService(queries, email)
		if err != nil {
			log.Printf("User %s is not an admin. Forbidden access.", email)
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to view this activity's registration"})
			return
		}
		log.Printf("User %s authorized as an admin.", email)
	}

	// Check if the member is the proposer of the activity

	// Fetch registered members for the activity
	members, err := services.GetSubmittedMembersService(queries, int32(activityID))
	if err != nil {
		log.Printf("Error fetching registered members for activity ID %d: %v", activityID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch registered members"})
		return
	}
	log.Printf("Fetched %d registered members for activity ID %d", len(members), activityID)

	// Return the registered members
	log.Printf("Returning registered members for activity ID %d", activityID)
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
