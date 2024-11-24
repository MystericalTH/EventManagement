package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"sinno-server/pkg/db"
	"sinno-server/pkg/services"

	"sinno-server/pkg/utils/typing"

	"github.com/gin-gonic/gin"
)

// Handler for getting all activities
func GetActivities(c *gin.Context, queries *db.Queries) {
	activityQueries, err := services.GetAllActivitiesService(queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities" + err.Error()})
		return
	}

	activities := []typing.Activity{}
	for _, query := range activityQueries {
		activity, err := typing.ConvertToActivity(db.ListActivityRow(query))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to format activity"})
			return
		}
		activities = append(activities, activity)
	}

	c.JSON(http.StatusOK, activities)
}

func GetProposerProposals(c *gin.Context, queries *db.Queries) {
	// Log the start of the handler
	log.Printf("Handler started: GetProposerProposals")

	// Retrieve the session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		log.Printf("Error retrieving session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
		return
	}
	log.Printf("Session retrieved successfully: %v", session.Values)

	// Extract user information and role from the session
	userInfo, userOk := session.Values["user"].(UserInfo)
	role, roleOk := session.Values["role"].(string)
	if !userOk || !roleOk {
		log.Printf("Authentication failed: userOk=%v, roleOk=%v", userOk, roleOk)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	log.Printf("User authenticated: email=%s, role=%s", userInfo.Email, role)

	email := userInfo.Email

	// Only allow proposers to fetch proposals
	if role != "member" {
		log.Printf("Access forbidden: user role is %s (expected 'member')", role)
		c.JSON(http.StatusForbidden, gin.H{"error": "Only proposers can fetch proposals"})
		return
	}

	// Get proposer ID from the service
	log.Printf("Fetching proposer ID for email: %s", email)
	proposerID, err := services.GetMemberIDByEmailService(queries, email) // Assuming "proposer" and "member" IDs are derived similarly
	if err != nil {
		log.Printf("Error fetching proposer ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get proposer ID"})
		return
	}
	log.Printf("Proposer ID fetched successfully: %d", proposerID)

	// Fetch proposer proposals using the service
	log.Printf("Fetching proposer proposals for proposerID: %d", proposerID)
	activityQueries, err := services.GetProposerProposalsService(queries, proposerID)
	if err != nil {
		log.Printf("Error fetching proposer proposals: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}
	log.Printf("Proposer proposals fetched successfully: %d activities found", len(activityQueries))

	// Convert query results to activities
	activities := []typing.Activity{}
	for _, query := range activityQueries {
		activity, err := typing.ConvertToActivity(db.ListActivityRow(query))
		if err != nil {
			log.Printf("Error converting query result to activity: %v", err)
			continue
		}
		activities = append(activities, activity)
	}
	log.Printf("Activities converted successfully: %d activities", len(activities))

	// Respond with the activities
	log.Printf("Responding with activities: %v", activities)
	c.JSON(http.StatusOK, activities)
}

// Handler for getting an activity by ID
func GetActivityByID(c *gin.Context, queries *db.Queries) {
	activityIDStr := c.Param("id")
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
	formattedActivity, err := typing.ConvertToActivity(activity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to ConvertToActivity"})
		return
	}
	c.JSON(http.StatusOK, formattedActivity)
}

// Handler for getting activity roles
func GetActivityRoles(c *gin.Context, queries *db.Queries) {
	activityIDStr := c.Param("id")
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
	Title          string   `json:"title" binding:"required"`
	Startdate      string   `json:"startDate" binding:"required"`
	Enddate        string   `json:"endDate" binding:"required"`
	Maxparticipant int32    `json:"maxParticipant" binding:"required"` // Updated to match frontend
	Format         string   `json:"format" binding:"required"`
	Description    string   `json:"description" binding:"required"`
	Advisor        *string  `json:"advisor"`
	Starttime      string   `json:"startTime"`
	Endtime        string   `json:"endTime"`
	Activityroles  []string `json:"activityRole" binding:"required"`
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
	if err != nil {
		log.Printf("PostActivity: Invalid end date: %s. Error: %v\n", req.Enddate, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	log.Printf("PostActivity: StartDate: %s, EndDate: %s validated successfully\n", startDate, endDate)

	// Ensure both start date and end date are in the future
	currentTime := time.Now()
	if !startDate.After(currentTime) || !endDate.After(currentTime) {
		log.Printf("PostActivity: StartDate or EndDate is not in the future. StartDate: %s, EndDate: %s, CurrentTime: %s\n", startDate, endDate, currentTime)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start date and end date must be in the future"})
		return
	}

	log.Printf("PostActivity: StartDate and EndDate are valid and in the future\n")

	// Validate and handle workshop-specific timing
	var startTime, endTime time.Time
	if req.Format == "workshop" {
		startTime, err = parseTime(req.Starttime)
		if err != nil {
			log.Printf("PostActivity: Invalid workshop start time: %s. Error: %v\n", req.Starttime, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workshop start time format"})
			return
		}
		endTime, err = parseTime(req.Endtime)
		if err != nil || !startTime.Before(endTime) {
			log.Printf("PostActivity: Invalid workshop end time: %s or start time not before end time: %v\n", req.Endtime, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Workshop end time must be after start time"})
			return
		}
		log.Printf("PostActivity: Workshop times validated successfully: StartTime: %s, EndTime: %s\n", startTime, endTime)
	}

	// Insert activity into the database
	// Insert activity into the database
	params := db.InsertActivityParams{
		Title:           req.Title,
		Proposer:        int32(memberID),
		Startdate:       startDate,
		Enddate:         endDate,
		Maxparticipant:  req.Maxparticipant,
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
		workshopParams := db.InsertWorkshopParams{
			Workshopid: activityID,
			Starttime:  startTime.Format("15:04"), // Convert time.Time to string in HH:MM format
			Endtime:    endTime.Format("15:04"),   // Convert time.Time to string in HH:MM format
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

	activityIDStr := c.Param("id")
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

	c.JSON(http.StatusNoContent, gin.H{"message": "Activity registration approved successfully"})
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
	activityQueries, err := services.GetAcceptedActivitiesService(queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}
	activities := []typing.Activity{}
	for _, query := range activityQueries {
		activity, _ := typing.ConvertToActivity(db.ListActivityRow(query))
		activities = append(activities, activity)
	}

	c.JSON(http.StatusOK, activities)
}
