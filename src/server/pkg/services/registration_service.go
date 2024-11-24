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

func CheckProposerService(queries *db.Queries, activityID, memberID int32) (bool, error) {
	result, err := queries.CheckProposer(context.Background(), db.CheckProposerParams{
		ActivityID: activityID,
		Proposer:   memberID,
	})
	if err != nil {
		return false, err
	}
	return result, nil
}

func GetSubmittedMembersService(queries *db.Queries, activityID int32) ([]db.ListSubmittedMembersRow, error) {
	return queries.ListSubmittedMembers(context.Background(), activityID)
}
