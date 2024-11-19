package services

import (
	"context"
	"sinno-server/pkg/db"
)

// Get all members service
func GetAllMembersService(queries *db.Queries) ([]db.ListAcceptedMembersRow, error) {
	return queries.ListAcceptedMembers(context.Background())
}

// Get member by ID service
func GetMemberByIDService(queries *db.Queries, memberID int32) (db.ListMemberRow, error) {
	return queries.ListMember(context.Background(), memberID)
}

// Create member service
func CreateMemberService(queries *db.Queries, params db.InsertMemberParams) error {
	return queries.InsertMember(context.Background(), params)
}

// Accept member service
func AcceptMemberService(queries *db.Queries, memberID int32) error {
	return queries.AcceptMember(context.Background(), memberID)
}

func GetAllMemberRequestsService(queries *db.Queries) ([]db.ListRequestingMembersRow, error) {
	return queries.ListRequestingMembers(context.Background())
}
