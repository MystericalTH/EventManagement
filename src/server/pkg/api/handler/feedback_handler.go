package handler

import (
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetFeedbackStatus handles retrieving feedback submission status
func GetFeedbackStatus(c *gin.Context, queries *db.Queries) {
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

	// Only allow members to check feedback status
	if role != "member" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only members can check feedback status"})
		return
	}

	// Get member ID from the service
	memberID, err := queries.GetMemberIDByEmail(c, email)
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

	// Check if the user has submitted feedback for this activity
	hasSubmitted, err := services.HasSubmittedFeedbackService(queries, int32(activityID), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check feedback status"})
		return
	}

	// Return JSON response
	c.JSON(http.StatusOK, gin.H{"hasSubmittedFeedback": hasSubmitted})
}

// SubmitFeedback handles submitting new feedback
func SubmitFeedback(c *gin.Context, queries *db.Queries) {
	// Ensure the request method is POST
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

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

	// Only allow members to submit feedback
	if role != "member" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only members can submit feedback"})
		return
	}

	memberID, err := queries.GetMemberIDByEmail(c, email)
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

	// Decode the request body
	var feedbackData db.Feedback
	if err := c.ShouldBindJSON(&feedbackData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create feedback parameters
	params := db.InsertFeedbackParams{
		Activityid:      int32(activityID),
		Memberid:        memberID,
		Feedbackmessage: feedbackData.Feedbackmessage,
	}

	// Insert the feedback using the service
	if err := services.CreateFeedbackService(queries, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit feedback"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Feedback submitted successfully"})
}
