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
	Startdate     string   `json:"startdate" binding:"required"`
	Enddate       string   `json:"enddate" binding:"required"`
	Maxnumber     int32    `json:"maxnumber" binding:"required"`
	Format        string   `json:"format" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	Advisor       *string  `json:"advisor,omitempty"`
	Starttime     string   `json:"starttime,omitempty"`
	Endtime       string   `json:"endtime,omitempty"`
	Activityroles []string `json:"activityRoles"`
}

// Handler for posting a new activity
func PostActivity(c *gin.Context, queries *db.Queries) {
	// Retrieve user info from session
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

	// Bind JSON request to CreateActivityRequest
	var req CreateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	startDate, err := parseDate(req.Startdate)
	if err != nil {
		log.Printf("StartDate parsing failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}
	endDate, err := parseDate(req.Enddate)
	if err != nil {
		log.Printf("EndDate parsing failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	// Validate that StartDate is before EndDate
	if !startDate.Before(endDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start date must be before end date"})
		return
	}

	// Set ProposeDateTime to current UTC time
	proposeDateTime := time.Now()

	// Prepare service params
	params := db.InsertActivityParams{
		Title:           req.Title,
		Proposer:        int32(memberID),
		Startdate:       startDate,
		Enddate:         endDate,
		Maxnumber:       req.Maxnumber,
		Format:          req.Format,
		Description:     req.Description,
		Proposedatetime: proposeDateTime,
	}

	// Call the service to create the activity
	if err := services.InsertActivityService(queries, params); err != nil {
		log.Printf("InsertActivityService failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activity"})
		return
	}

	// Retrieve the activity ID using the title
	activityID, err := services.GetActivityIDByTitleService(queries, req.Title)
	if err != nil {
		log.Printf("GetActivityIDByTitleService failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activity ID"})
		return
	}

	if req.Format == "project" {
		// Insert project
		if req.Advisor != nil {
			projectParams := db.InsertProjectParams{
				Projectid: activityID,
				Advisor:   sql.NullString{String: *req.Advisor, Valid: req.Advisor != nil},
			}
			if err := services.InsertProjectService(queries, projectParams); err != nil {
				log.Printf("InsertProjectService failed: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert project"})
				return
			}
		}
	} else if req.Format == "workshop" {
		// Insert workshop
		if req.Starttime != "" && req.Endtime != "" {
			startTime, err := parseTime(req.Starttime)
			if err != nil {
				log.Printf("StartTime parsing failed: %v", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
				return
			}
			endTime, err := parseTime(req.Endtime)
			if err != nil {
				log.Printf("EndTime parsing failed: %v", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format"})
				return
			}
			workshopParams := db.InsertWorkshopParams{
				Workshopid: activityID,
				Starttime:  startTime,
				Endtime:    endTime,
			}
			if err := services.InsertWorkshopService(queries, workshopParams); err != nil {
				log.Printf("InsertWorkshopService failed: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert workshop"})
				return
			}
		}
	}

	// Insert activity roles
	for _, role := range req.Activityroles {
		roleParams := db.InsertActivityRoleParams{
			Activityid:   activityID,
			Activityrole: role,
		}
		if err := services.InsertActivityRoleService(queries, roleParams); err != nil {
			log.Printf("InsertActivityRoleService failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert activity role"})
			return
		}
	}

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
