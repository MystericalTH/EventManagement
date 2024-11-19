package services

import (
	"context"
	"sinno-server/pkg/db"
)

// Get all activities service
func GetAllActivitiesService(queries *db.Queries) ([]db.Activity, error) {
	return queries.ListRequestingActivities(context.Background())
}

// Get activity by ID service
func GetActivityByIDService(queries *db.Queries, activityID int32) (db.Activity, error) {
	return queries.ListActivity(context.Background(), activityID)
}

// Create activity service
func CreateActivityService(queries *db.Queries, params db.InsertActivityParams) error {
	return queries.InsertActivity(context.Background(), params)
}

// Get activity roles service
func GetActivityRolesService(queries *db.Queries, activityID int32) ([]string, error) {
	return queries.ListActivityRoles(context.Background(), activityID)
}

// Insert activity role service
func InsertActivityRoleService(queries *db.Queries, activityID int32, role string) error {
	params := db.InsertActivityRoleParams{
		Activityid:   activityID,
		Activityrole: role,
	}
	return queries.InsertActivityRole(context.Background(), params)
}
