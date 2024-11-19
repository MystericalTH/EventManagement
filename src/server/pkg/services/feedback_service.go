package services

import (
	"context"
	"sinno-server/pkg/db"
)

// CreateFeedbackService creates a new feedback entry
func CreateFeedbackService(queries *db.Queries, params db.CreateFeedbackParams) error {
	return queries.CreateFeedback(context.Background(), params)
}

// GetFeedbackByIDService retrieves a feedback entry by ID
func GetFeedbackByIDService(queries *db.Queries, feedbackID int32) (db.Feedback, error) {
	return queries.GetFeedbackByID(context.Background(), feedbackID)
}

// ListFeedbacksService lists all feedback entries
func GetAllFeedbacksService(queries *db.Queries) ([]db.Feedback, error) {
	return queries.ListFeedbacks(context.Background())
}

// HasSubmittedFeedbackService checks if feedback has been submitted by a member for an activity
func HasSubmittedFeedbackService(queries *db.Queries, activityID int32, memberID int32) (bool, error) {
	params := db.HasSubmittedFeedbackParams{
		Activityid: activityID,
		Memberid:   memberID,
	}
	return queries.HasSubmittedFeedback(context.Background(), params)
}
