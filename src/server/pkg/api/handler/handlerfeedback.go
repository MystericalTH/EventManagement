package api

import (
	"encoding/json"
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/models"
	"strconv"

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

	// Check if the user has submitted feedback for this activity
	var count int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM Feedback WHERE activity_id = ? AND member_id = ?", activityID, memberID).Scan(&count)
	if err != nil {
		http.Error(w, "Failed to check feedback status", http.StatusInternalServerError)
		return
	}
	hasSubmittedFeedback := count > 0

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"hasSubmittedFeedback": hasSubmittedFeedback})
}
