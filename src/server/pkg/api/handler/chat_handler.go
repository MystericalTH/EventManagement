package handler

import (
	"database/sql"
	"log"
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/services"

	"github.com/gin-gonic/gin"
)

// Create a new chat entry
func CreateChat(c *gin.Context, queries *db.Queries) {
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

	// Decode the request body
	var chatData db.InsertChatParams
	if err := c.ShouldBindJSON(&chatData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// Set adminID or developerID based on the role
	if role == "admin" {
		adminID, err := services.GetAdminIDByEmailService(queries, email)
		if err != nil {
			log.Printf("PostActivity: Failed to fetch admin ID for email %s: %v\n", email, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch admin ID", "details": err.Error()})
			return
		}
		chatData.Adminid = sql.NullInt32{Int32: int32(adminID), Valid: true}
		chatData.Developerid = sql.NullInt32{Valid: false}
	} else if role == "developer" {
		developerID, err := services.GetDeveloperIDByEmailService(queries, email)
		if err != nil {
			log.Printf("PostActivity: Failed to fetch admin ID for email %s: %v\n", email, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch admin ID", "details": err.Error()})
			return
		}
		chatData.Developerid = sql.NullInt32{Int32: int32(developerID), Valid: true}
		chatData.Adminid = sql.NullInt32{Valid: false}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Members cannot submit chats"})
		return
	}

	// Insert the chat using the service
	if err := services.CreateChatService(queries, chatData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit chat", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Chat submitted successfully"})
}

// Get all chats
func GetChats(c *gin.Context, queries *db.Queries) {
	// Retrieve the session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session", "details": err.Error()})
		return
	}

	role, roleOk := session.Values["role"].(string)
	if !roleOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if role == "member" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Members cannot access this page"})
		return
	}

	chats, err := services.GetChatsService(queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get feedback entries", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"chats": chats})
}
