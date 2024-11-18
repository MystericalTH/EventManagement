package api

import (
	"encoding/json"
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetRegistrationStatus(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session
	session, err := sessionStore.Get(r, sessionName)
	if err != nil {
		http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
		return
	}

	// Get user information and role from session
	userInfo, userOk := session.Values["user"].(models.UserInfo)
	role, roleOk := session.Values["role"].(string)
	if !userOk || !roleOk {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	email := userInfo.Email

	// Only allow members to submit feedback
	if role != "member" {
		http.Error(w, "Only members can submit feedback", http.StatusForbidden)
		return
	}

	// Get member ID from the Member table
	var memberID int64
	err = db.DB.QueryRow("SELECT member_id FROM Member WHERE email = ?", email).Scan(&memberID)
	if err != nil {
		http.Error(w, "Failed to get member ID", http.StatusInternalServerError)
		return
	}

	// Get activity ID from URL
	vars := mux.Vars(r)
	activityIDStr := vars["activityId"]
	activityID, err := strconv.ParseInt(activityIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid activity ID", http.StatusBadRequest)
		return
	}

	// Check if the user has registered for this activity
	var count int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM Registration WHERE activityID = ? AND memberID = ?", activityID, memberID).Scan(&count)
	if err != nil {
		http.Error(w, "Failed to check registration status", http.StatusInternalServerError)
		return
	}
	isRegistered := count > 0

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"isRegistered": isRegistered})
}

func SubmitRegistration(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session
	session, err := sessionStore.Get(r, sessionName)
	if err != nil {
		http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
		return
	}

	// Get user information and role from session
	userInfo, userOk := session.Values["user"].(models.UserInfo)
	role, roleOk := session.Values["role"].(string)
	if !userOk || !roleOk {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	email := userInfo.Email

	// Only allow members to register
	if role != "member" {
		http.Error(w, "Only members can register", http.StatusForbidden)
		return
	}

	// Get member ID from the Member table
	var memberID int64
	err = db.DB.QueryRow("SELECT memberID FROM Member WHERE email = ?", email).Scan(&memberID)
	if err != nil {
		http.Error(w, "Failed to get member ID", http.StatusInternalServerError)
		return
	}

	var registrationData db.Pjregist
	if err := json.NewDecoder(r.Body).Decode(&registrationData); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Insert the registration details into the Registration table
	_, err = db.DB.Exec("INSERT INTO Registration (activityID, memberID, role, expectation, datetime) VALUES (?, ?, ?, ?, ?)",
		registrationData.Projectid,
		memberID,
		registrationData.Role,
		registrationData.Expectation,
		time.Now().Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		http.Error(w, "Failed to submit registration", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registration submitted successfully"})
}
