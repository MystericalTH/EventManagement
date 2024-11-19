package handler

import (
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler for getting all members
func GetAllMembers(c *gin.Context, queries *db.Queries) {
	members, err := services.GetAllMembersService(queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, members)
}

// Handler for getting a member by ID
func GetMemberByID(c *gin.Context, queries *db.Queries) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
		return
	}

	member, err := services.GetMemberByIDService(queries, int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}
	c.JSON(http.StatusOK, member)
}

// Handler for creating a new member
func CreateMember(c *gin.Context, queries *db.Queries) {
	var params db.InsertMemberParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := services.CreateMemberService(queries, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Member created successfully"})
}

// Handler for accepting a member
func AcceptMember(c *gin.Context, queries *db.Queries) {
	// Retrieve user email from session
	session, err := sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.String(http.StatusUnauthorized, "Session retrieval failed")
		return
	}

	userEmail, ok := session.Values["user_email"].(string)
	if !ok || userEmail == "" {
		c.String(http.StatusUnauthorized, "Unauthorized: User email not found in session")
		return
	}

	// Fetch adminID using the email
	adminID, err := services.FetchAdminIDService(queries, userEmail)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch admin ID: %s", err.Error())
		return
	}

	// Parse memberID from URL parameters
	memberIDParam := c.Param("id")
	memberID, err := strconv.Atoi(memberIDParam)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid member ID: %s", err.Error())
		return
	}

	// Call service to accept the member
	err = services.AcceptMemberService(queries, int32(memberID), adminID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to accept member: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member accepted successfully"})
}

func GetAllMemberRequests(c *gin.Context, queries *db.Queries) {
	memberRequests, err := services.GetAllMemberRequestsService(queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, memberRequests)
}
