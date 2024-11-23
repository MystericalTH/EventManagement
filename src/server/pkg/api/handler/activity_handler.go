package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"sinno-server/pkg/db"
	"sinno-server/pkg/services"

	"github.com/gin-gonic/gin"
)

// Handler for getting all activities
func GetActivities(c *gin.Context, queries *db.Queries) {
	activities, err := services.GetAllActivitiesService(queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// Handler for getting an activity by ID
func GetActivityByID(c *gin.Context, queries *db.Queries) {
	activityIDStr := c.Param("activityId")
	activityID, err := strconv.Atoi(activityIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	activity, err := services.GetActivityByIDService(queries, int32(activityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activity"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// Handler for getting activity roles
func GetActivityRoles(c *gin.Context, queries *db.Queries) {
	activityIDStr := c.Param("activityId")
	activityID, err := strconv.Atoi(activityIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	roles, err := services.GetActivityRolesService(queries, int32(activityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activity roles"})
		return
	}

	c.JSON(http.StatusOK, roles)
}

type CreateActivityRequest struct {
	Title         string   `json:"title" binding:"required"`
	Startdate     string   `json:"startDate" binding:"required"`
	Enddate       string   `json:"endDate" binding:"required"`
	Maxnumber     int32    `json:"maxParticipant" binding:"required"` // Updated to match frontend
	Format        string   `json:"format" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	Advisor       *string  `json:"advisor"`
	Starttime     string   `json:"startTime"`
	Endtime       string   `json:"endTime"`
	Activityroles []string `json:"roles" binding:"required"`
}

func PostActivity(c *gin.Context, queries *db.Queries) {
	log.Println("PostActivity: Started handling request")

	// Retrieve user email and role from session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		log.Printf("PostActivity: Failed to retrieve session: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session retrieval failed"})
		return
	}

	userEmail, emailOk := session.Values["user_email"].(string)
	role, roleOk := session.Values["role"].(string)
	if !emailOk || !roleOk || role != "member" {
		log.Printf("PostActivity: Unauthorized access. EmailOk: %v, RoleOk: %v, Role: %v\n", emailOk, roleOk, role)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid session or insufficient privileges"})
		return
	}

	log.Printf("PostActivity: Retrieved session for email: %s and role: %s\n", userEmail, role)

	// Fetch member ID using the email
	memberID, err := services.GetMemberIDByEmailService(queries, userEmail)
	if err != nil {
		log.Printf("PostActivity: Failed to fetch member ID for email %s: %v\n", userEmail, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch member ID", "details": err.Error()})
		return
	}

	log.Printf("PostActivity: Retrieved member ID: %d for email: %s\n", memberID, userEmail)

	// Parse and validate request body
	var req CreateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("PostActivity: Failed to bind JSON request: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	log.Printf("PostActivity: Request payload: %+v\n", req)

	// Validate and parse date fields
	startDate, err := parseDate(req.Startdate)
	if err != nil {
		log.Printf("PostActivity: Invalid start date: %s. Error: %v\n", req.Startdate, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}
	endDate, err := parseDate(req.Enddate)
	if err != nil || !startDate.Before(endDate) {
		log.Printf("PostActivity: Invalid end date: %s or start date not before end date: %v\n", req.Enddate, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "End date must be after start date"})
		return
	}

	log.Printf("PostActivity: StartDate: %s, EndDate: %s validated successfully\n", startDate, endDate)

	// Insert activity into the database
	params := db.InsertActivityParams{
		Title:           req.Title,
		Proposer:        int32(memberID),
		Startdate:       startDate,
		Enddate:         endDate,
		Maxnumber:       req.Maxnumber,
		Format:          req.Format,
		Description:     req.Description,
		Proposedatetime: time.Now(),
	}
	log.Printf("PostActivity: InsertActivityParams: %+v\n", params)

	err = services.InsertActivityService(queries, params)
	if err != nil {
		log.Printf("PostActivity: Failed to insert activity: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activity", "details": err.Error()})
		return
	}

	// Retrieve the activity ID
	activityID, err := services.GetActivityIDByTitleService(queries, req.Title)
	if err != nil {
		log.Printf("PostActivity: Failed to retrieve activity ID for title %s: %v\n", req.Title, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activity ID", "details": err.Error()})
		return
	}

	log.Printf("PostActivity: Retrieved activity ID: %d for title: %s\n", activityID, req.Title)

	// Handle specific formats (project or workshop)
	if req.Format == "project" && req.Advisor != nil {
		projectParams := db.InsertProjectParams{
			Projectid: activityID,
			Advisor:   sql.NullString{String: *req.Advisor, Valid: true},
		}
		log.Printf("PostActivity: InsertProjectParams: %+v\n", projectParams)
		if err := services.InsertProjectService(queries, projectParams); err != nil {
			log.Printf("PostActivity: Failed to insert project: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert project", "details": err.Error()})
			return
		}
	} else if req.Format == "workshop" {
		startTime, startErr := parseTime(req.Starttime)
		endTime, endErr := parseTime(req.Endtime)
		if startErr != nil || endErr != nil {
			log.Printf("PostActivity: Invalid workshop times. StartTime: %s, EndTime: %s. Errors: %v, %v\n", req.Starttime, req.Endtime, startErr, endErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workshop times"})
			return
		}
		workshopParams := db.InsertWorkshopParams{
			Workshopid: activityID,
			Starttime:  startTime,
			Endtime:    endTime,
		}
		log.Printf("PostActivity: InsertWorkshopParams: %+v\n", workshopParams)
		if err := services.InsertWorkshopService(queries, workshopParams); err != nil {
			log.Printf("PostActivity: Failed to insert workshop: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert workshop", "details": err.Error()})
			return
		}
	}

	// Insert activity roles
	for _, role := range req.Activityroles {
		roleParams := db.InsertActivityRoleParams{
			Activityid:   activityID,
			Activityrole: role,
		}
		log.Printf("PostActivity: InsertActivityRoleParams: %+v\n", roleParams)
		if err := services.InsertActivityRoleService(queries, roleParams); err != nil {
			log.Printf("PostActivity: Failed to insert activity roles: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert activity roles", "details": err.Error()})
			return
		}
	}

	// Respond with success
	log.Println("PostActivity: Activity created successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Activity created successfully"})
}

// Helper function to parse date
func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// Helper function to parse time
func parseTime(timeStr string) (time.Time, error) {
	return time.Parse("15:04", timeStr)
}

func ApproveActivityRegistration(c *gin.Context, queries *db.Queries) {
	// Retrieve user info from session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
		return
	}
	userEmail, ok := session.Values["user_email"].(string)
	if !ok || userEmail == "" {
		c.String(http.StatusUnauthorized, "Unauthorized: User email not found in session")
		return
	}

	adminID, err := services.GetAdminIDByEmailService(queries, userEmail)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch admin ID: %s", err.Error())
		return
	}

	activityIDStr := c.Param("activityId")
	activityID, err := strconv.Atoi(activityIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	// Call the service to approve the activity registration
	if err := services.ApproveActivityRegistrationService(queries, int32(activityID), adminID); err != nil {
		log.Printf("ApproveActivityRegistration failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve activity registration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity registration approved successfully"})
}

func DeleteActivity(c *gin.Context, queries *db.Queries) {
	ActivitityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
		return
	}

	err = services.DeleteActivityService(queries, int32(ActivitityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Activity deleted successfully"})
}

func GetAcceptedActivities(c *gin.Context, queries *db.Queries) {
	activities, err := services.GetAcceptedActivitiesService(queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}
	c.JSON(http.StatusOK, activities)
}
