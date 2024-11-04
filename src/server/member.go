package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Member struct {
	MemberID       int       `json:"memberID"`
	FName          string    `json:"fName"`
	LName          string    `json:"lName"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	GithubUrl      string    `json:"githubUrl"`
	Interest       string    `json:"interest"`
	Reason         string    `json:"reason"`
	AcceptDateTime time.Time `json:"acceptDateTime"`
	AcceptAdmin    int       `json:"acceptAdmin"`
}

func memberHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMember(w, r)
	case "POST":
		postMember(w, r)
	case "PATCH":
		patchMember(w, r)
	case "DELETE":
		deleteMember(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getMember(w http.ResponseWriter, r *http.Request) {
	// Extract member ID from URL query parameters
	memberID := r.URL.Query().Get("memberID")
	if memberID == "" {
		http.Error(w, "memberID not provided", http.StatusBadRequest)
		return
	}

	// Query the database for member information
	var member Member
	err := db.QueryRow("SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason, acceptDateTime, acceptAdmin FROM members WHERE memberID = ?", memberID).
		Scan(&member.MemberID, &member.FName, &member.LName, &member.Email, &member.Phone, &member.GithubUrl, &member.Interest, &member.Reason, &member.AcceptDateTime, &member.AcceptAdmin)
	if err != nil {
		http.Error(w, "Failed to fetch member information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return member information as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(member)
}

func postMember(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request body into a Member struct
	var member Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Insert member into the database
	_, err := db.Exec("INSERT INTO members (fName, lName, email, phone, githubUrl, interest, reason, acceptDateTime, acceptAdmin) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		member.FName, member.LName, member.Email, member.Phone, member.GithubUrl, member.Interest, member.Reason, member.AcceptDateTime, member.AcceptAdmin)
	if err != nil {
		http.Error(w, "Failed to insert member: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Member created successfully"))
}

func patchMember(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request body into a Member struct
	var member Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate memberID
	if member.MemberID == 0 {
		http.Error(w, "memberID is required", http.StatusBadRequest)
		return
	}

	// Build the SQL query dynamically based on provided fields
	query := "UPDATE members SET "
	params := []interface{}{}

	if member.FName != "" {
		query += "fName = ?, "
		params = append(params, member.FName)
	}
	if member.LName != "" {
		query += "lName = ?, "
		params = append(params, member.LName)
	}
	if member.Email != "" {
		query += "email = ?, "
		params = append(params, member.Email)
	}
	if member.Phone != "" {
		query += "phone = ?, "
		params = append(params, member.Phone)
	}
	if member.GithubUrl != "" {
		query += "githubUrl = ?, "
		params = append(params, member.GithubUrl)
	}
	if member.Interest != "" {
		query += "interest = ?, "
		params = append(params, member.Interest)
	}
	if member.Reason != "" {
		query += "reason = ?, "
		params = append(params, member.Reason)
	}

	// Remove the trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE memberID = ?"
	params = append(params, member.MemberID)

	// Update member in the database
	_, err := db.Exec(query, params...)
	if err != nil {
		http.Error(w, "Failed to update member: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Write([]byte("Member updated successfully"))
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	// Extract member ID from the request
	memberID := r.URL.Query().Get("memberID")
	if memberID == "" {
		http.Error(w, "memberID not provided", http.StatusBadRequest)
		return
	}

	// Delete member from the database
	_, err := db.Exec("DELETE FROM members WHERE memberID = ?", memberID)
	if err != nil {
		http.Error(w, "Failed to delete member: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Write([]byte("Member deleted successfully"))
}

func member() {
	// Initialize the database connection (replace with your DSN)
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/activitymanagement")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	http.HandleFunc("/member", memberHandler)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
