package services

import (
	"context"
	"database/sql"
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
func AcceptMemberService(queries *db.Queries, memberID int32, adminID int32) error {
	params := db.AcceptMemberParams{
		Acceptadmin: sql.NullInt32{Int32: adminID, Valid: true},
		Memberid:    memberID,
	}
	return queries.AcceptMember(context.Background(), params)
}

func GetAllMemberRequestsService(queries *db.Queries) ([]db.ListRequestingMembersRow, error) {
	return queries.ListRequestingMembers(context.Background())
}

func UpdateMemberService(queries *db.Queries, params db.UpdateMemberParams) error {
	return queries.UpdateMember(context.Background(), params)
}

func DeleteMemberService(queries *db.Queries, memberID int32) error {
	return queries.DeleteMember(context.Background(), memberID)
}
