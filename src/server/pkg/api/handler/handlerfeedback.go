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

func GetFeedbackStatus(w http.ResponseWriter, r *http.Request) {
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
	err = db.DB.QueryRow("SELECT memberID FROM Member WHERE email = ?", email).Scan(&memberID)
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

	// Check if the user has submitted feedback for this activity
	var count int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM Feedback WHERE activityID = ? AND memberID = ?", activityID, memberID).Scan(&count)
	if err != nil {
		http.Error(w, "Failed to check feedback status", http.StatusInternalServerError)
		return
	}
	hasSubmittedFeedback := count > 0

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"hasSubmittedFeedback": hasSubmittedFeedback})
}

func SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
	err = db.DB.QueryRow("SELECT memberID FROM Member WHERE email = ?", email).Scan(&memberID)
	if err != nil {
		http.Error(w, "Failed to get member ID", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	activityIDStr := vars["activityId"]
	activityID, err := strconv.ParseInt(activityIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid activity ID", http.StatusBadRequest)
		return
	}

	// Decode the request body
	var feedbackData db.Feedback
	err = json.NewDecoder(r.Body).Decode(&feedbackData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if feedback already exists
	var count int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM Feedback WHERE activityID = ? AND memberID = ?", activityID, memberID).Scan(&count)
	if err != nil {
		http.Error(w, "Failed to check existing feedback", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Feedback already submitted", http.StatusConflict)
		return
	}

	// Insert the feedback into the database
	_, err = db.DB.Exec("INSERT INTO Feedback (activityID, memberID, feedbackMessage, feedbackDateTime) VALUES (?, ?, ?, ?)",
		activityID,
		memberID,
		feedbackData.Feedbackmessage,
		time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		http.Error(w, "Failed to submit feedback", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
