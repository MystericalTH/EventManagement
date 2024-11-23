package services

import (
	"context"
	"sinno-server/pkg/db"
)

// CreateFeedbackService creates a new feedback entry
func CreateFeedbackService(queries *db.Queries, params db.InsertFeedbackParams) error {
	return queries.InsertFeedback(context.Background(), params)
}

// ListFeedbacksByActivityService lists all feedback entries for an activity
func GetFeedbacksByActivityService(queries *db.Queries, activityID int32) ([]db.ListFeedbacksRow, error) {
	return queries.ListFeedbacks(context.Background(), activityID)
}

// HasSubmittedFeedbackService checks if feedback has been submitted by a member for an activity
func HasSubmittedFeedbackService(queries *db.Queries, activityID int32, memberID int32) (bool, error) {
	params := db.HasSubmittedFeedbackParams{
		Activityid: activityID,
		Memberid:   memberID,
	}
	return queries.HasSubmittedFeedback(context.Background(), params)
}
