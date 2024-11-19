package handler

import (
	"net/http"
	"strconv"

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

type CreateActivityRequest struct {
	db.InsertActivityParams
	Activityroles []string `json:"activityRoles"`
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

// Handler for posting a new activity
