package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sinno-server/pkg/db"
	secure "sinno-server/pkg/utils/dbsecure"
)

// Get all members service
func GetAllMembersService(queries *db.Queries) ([]db.ListAcceptedMembersRow, error) {
	members, err := queries.ListAcceptedMembers(context.Background())
	if err != nil {
		return nil, err
	}

	for i, member := range members {
		// Decrypt first name
		decryptedFName, err := secure.DecryptFromString(member.Fname, os.Getenv("ENCRYPT_KEY"))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt first name for member ID %d: %v", member.Memberid, err)
		}

		// Decrypt last name
		decryptedLName, err := secure.DecryptFromString(member.Lname, os.Getenv("ENCRYPT_KEY"))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt last name for member ID %d: %v", member.Memberid, err)
		}

		// Decrypt phone
		decryptedPhone, err := secure.DecryptFromString(member.Phone, os.Getenv("ENCRYPT_KEY"))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt phone for member ID %d: %v", member.Memberid, err)
		}

		// Update member with decrypted values
		members[i].Fname = decryptedFName
		members[i].Lname = decryptedLName
		members[i].Phone = decryptedPhone
	}

	return members, nil
}

// Get a specific member by ID
func GetMemberByIDService(queries *db.Queries, memberID int32) (db.ListMemberRow, error) {
	member, err := queries.ListMember(context.Background(), memberID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Member with ID %d not found", memberID)
			return member, fmt.Errorf("member with ID %d not found", memberID)
		}
		log.Printf("Database error while fetching member ID %d: %v", memberID, err)
		return member, err
	}

	// Decrypt fields
	member, err = decryptMemberFields(member)
	if err != nil {
		return member, err
	}

	log.Printf("Successfully retrieved and decrypted member ID %d: %+v", memberID, member)
	return member, nil
}

// Create a new member service
func CreateMemberService(queries *db.Queries, member db.InsertMemberParams) error {
	log.Printf("Starting CreateMemberService for member: %+v", member)

	// Encrypt fields
	encryptedFName, err := secure.EncryptToString(member.Fname, os.Getenv("ENCRYPT_KEY"))
	if err != nil {
		return fmt.Errorf("failed to encrypt first name: %v", err)
	}
	encryptedLName, err := secure.EncryptToString(member.Lname, os.Getenv("ENCRYPT_KEY"))
	if err != nil {
		return fmt.Errorf("failed to encrypt last name: %v", err)
	}
	encryptedPhone, err := secure.EncryptToString(member.Phone, os.Getenv("ENCRYPT_KEY"))
	if err != nil {
		return fmt.Errorf("failed to encrypt phone: %v", err)
	}

	// Update member with encrypted fields
	member.Fname = encryptedFName
	member.Lname = encryptedLName
	member.Phone = encryptedPhone

	err = queries.InsertMember(context.Background(), member)
	if err != nil {
		return fmt.Errorf("error inserting member into database: %v", err)
	}

	log.Printf("Member successfully created: %+v", member)
	return nil
}

// Accept member service
func AcceptMemberService(queries *db.Queries, memberID int32, adminID int32) error {
	params := db.AcceptMemberParams{
		Acceptadmin: sql.NullInt32{Int32: adminID, Valid: true},
		Memberid:    memberID,
	}
	return queries.AcceptMember(context.Background(), params)
}

// Get all member requests service
func GetAllMemberRequestsService(queries *db.Queries) ([]db.ListRequestingMembersRow, error) {
	members, err := queries.ListRequestingMembers(context.Background())
	if err != nil {
		log.Printf("Error fetching member requests: %v", err)
		return nil, err
	}

	for i, member := range members {
		// Decrypt fields
		decryptedFName, err := secure.DecryptFromString(member.Fname, os.Getenv("ENCRYPT_KEY"))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt first name for member ID %d: %v", member.Memberid, err)
		}
		decryptedLName, err := secure.DecryptFromString(member.Lname, os.Getenv("ENCRYPT_KEY"))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt last name for member ID %d: %v", member.Memberid, err)
		}
		decryptedPhone, err := secure.DecryptFromString(member.Phone, os.Getenv("ENCRYPT_KEY"))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt phone for member ID %d: %v", member.Memberid, err)
		}

		// Update member with decrypted values
		members[i].Fname = decryptedFName
		members[i].Lname = decryptedLName
		members[i].Phone = decryptedPhone
	}

	log.Printf("Successfully decrypted %d member requests", len(members))
	return members, nil
}

// Delete member service
func DeleteMemberService(queries *db.Queries, memberID int32) error {
	return queries.DeleteMember(context.Background(), memberID)
}

// Helper function to decrypt member fields
func decryptMemberFields(member db.ListMemberRow) (db.ListMemberRow, error) {
	decryptedFName, err := secure.DecryptFromString(member.Fname, os.Getenv("ENCRYPT_KEY"))
	if err != nil {
		return member, fmt.Errorf("failed to decrypt first name: %v", err)
	}

	decryptedLName, err := secure.DecryptFromString(member.Lname, os.Getenv("ENCRYPT_KEY"))
	if err != nil {
		return member, fmt.Errorf("failed to decrypt last name: %v", err)
	}

	decryptedPhone, err := secure.DecryptFromString(member.Phone, os.Getenv("ENCRYPT_KEY"))
	if err != nil {
		return member, fmt.Errorf("failed to decrypt phone: %v", err)
	}

	member.Fname = decryptedFName
	member.Lname = decryptedLName
	member.Phone = decryptedPhone

	return member, nil
}
