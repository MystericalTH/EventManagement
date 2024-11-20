package services

import (
	"context"
	"sinno-server/pkg/db"
)

// CreateRegistrationService creates a new registration entry
func CreateRegistrationService(queries *db.Queries, params db.InsertRegistrationParams) error {
	return queries.InsertRegistration(context.Background(), params)
}

// GetRegistrationStatusService checks if a member is registered for an activity
func GetRegistrationStatusService(queries *db.Queries, activityID int32, memberID int32) (bool, error) {
	params := db.GetRegistrationStatusParams{
		Activityid: activityID,
		Memberid:   memberID,
	}
	return queries.GetRegistrationStatus(context.Background(), params)
}
