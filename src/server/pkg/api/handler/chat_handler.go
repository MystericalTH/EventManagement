package handler

import (
	"log"
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"sinno-server/pkg/utils/typing"
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
	var chatData db.InsertAdminDevChatParams
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
		chatData.Sender = "admin"
		chatData.Adminid = adminID
	} else if role == "developer" {
		developerID, err := services.GetDeveloperIDByEmailService(queries, email)
		if err != nil {
			log.Printf("PostActivity: Failed to fetch admin ID for email %s: %v\n", email, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch admin ID", "details": err.Error()})
			return
		}
		chatData.Sender = "developer"
		chatData.Developerid = developerID
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Members cannot submit chats"})
		return
	}

	// Insert the chat using the service
	if err := services.InsertAdminDevChatService(queries, chatData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit chat", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Chat submitted successfully"})
}

// Get all chats
func ListAdminDevChats(c *gin.Context, queries *db.Queries) {
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
	userInfo, userInfoOk := session.Values["user"].(UserInfo)
	if !userInfoOk {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse userInfo"})
		return
	}

	params := db.ListAdminDevChatParams{}

	if role == "developer" {
		params.Developerid, err = services.GetDeveloperIDByEmailService(queries, userInfo.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get email"})
			return
		}
		adminIDStr := c.Param("id")
		adminID, err := strconv.ParseInt(adminIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing admin ID"})
			return
		}
		params.Adminid = int32(adminID)
	} else if role == "admin" {
		params.Adminid, err = services.GetAdminIDByEmailService(queries, userInfo.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get email"})
			return
		}
		developerIDStr := c.Param("id")
		developerID, err := strconv.ParseInt(developerIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing admin ID"})
			return
		}
		params.Developerid = int32(developerID)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "You cannot access this page"})
		return
	}

	chats, err := services.ListAdminDevChatService(queries, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get feedback entries", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"chats": chats})
}

func ListInitialAdminDevChat(c *gin.Context, queries *db.Queries) {
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot get session"})
		return
	}

	role, roleOk := session.Values["role"].(string)
	if !roleOk {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userInfo, userInfoOk := session.Values["user"].(UserInfo)
	if !userInfoOk {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse userInfo"})
		return
	}

	if role == "admin" {
		adminID, err := queries.GetAdminIDByEmail(c, userInfo.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get admin email"})
			return
		}
		data, err := services.ListInitialAdminChatToDevService(queries, adminID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot retrieve chat", "details": err.Error()})
			return
		}
		var response = []typing.ChatChannelInfo{}
		for _, d := range data {
			response = append(response, typing.ConvertListInitialAdminChatToDevRow(d))
		}
		c.JSON(http.StatusOK, gin.H{"data": response})
	} else if role == "developer" {
		developerID, err := queries.GetDeveloperIDByEmail(c, userInfo.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get developer email"})
			return
		}
		data, err := services.ListInitialDevChatToAdminService(queries, developerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot retrieve chat", "details": err.Error()})
			return
		}
		var response = []typing.ChatChannelInfo{}
		for _, d := range data {
			response = append(response, typing.ConvertListInitialDevChatToAdminRow(d))
		}
		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}
