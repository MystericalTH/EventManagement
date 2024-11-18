package api

import (
	"encoding/json"
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/models"
	"strconv"
)

func GetActivities(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT * FROM Activity")
	if err != nil {
		http.Error(w, "Failed to fetch activities", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var activities []models.Activity
	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(
			&activity.ActivityID,
			&activity.Title,
			&activity.Proposer,
			&activity.StartDate,
			&activity.EndDate,
			&activity.MaxNumber,
			&activity.Format,
			&activity.Description,
			&activity.ProposeDateTime,
			&activity.AcceptAdmin,
			&activity.AcceptDateTime,
			&activity.ApplicationStatus,
		); err != nil {
			http.Error(w, "Failed to scan activity", http.StatusInternalServerError)
			return
		}
		activities = append(activities, activity)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}

func GetActivityByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid activity ID", http.StatusBadRequest)
		return
	}

	var activity models.Activity
	err = db.DB.QueryRow("SELECT * FROM Activity WHERE activity_id = ?", id).Scan(
		&activity.ActivityID,
		&activity.Title,
		&activity.Proposer,
		&activity.StartDate,
		&activity.EndDate,
		&activity.MaxNumber,
		&activity.Format,
		&activity.Description,
		&activity.ProposeDateTime,
		&activity.AcceptAdmin,
		&activity.AcceptDateTime,
		&activity.ApplicationStatus,
	)
	if err != nil {
		http.Error(w, "Failed to fetch activity", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
}

func PostActivity(w http.ResponseWriter, r *http.Request) {
	// Retrieve user info from session
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

	var proposerID int32

	if role == "admin" {
		// Look up in Admin table
		err = db.DB.QueryRow("SELECT admin_id FROM Admin WHERE email = ?", email).Scan(&proposerID)
		if err != nil {
			http.Error(w, "Failed to get admin ID", http.StatusInternalServerError)
			return
		}
	} else if role == "member" {
		// Look up in Member table
		err = db.DB.QueryRow("SELECT member_id FROM Member WHERE email = ?", email).Scan(&proposerID)
		if err != nil {
			http.Error(w, "Failed to get member ID", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unknown user role", http.StatusUnauthorized)
		return
	}

	var activity struct {
		models.Activity
		ActivityRoles []string `json:"activityRole"`
	}

	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "Failed to decode activity", http.StatusInternalServerError)
		return
	}

	activity.Proposer = proposerID

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	result, err := tx.Exec("INSERT INTO Activity (title, proposer, start_date, end_date, max_number, format, description, propose_date_time, application_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		activity.Title,
		activity.Proposer,
		activity.StartDate,
		activity.EndDate,
		activity.MaxNumber,
		activity.Format,
		activity.Description,
		activity.ProposeDateTime,
		activity.ApplicationStatus,
	)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to insert activity", http.StatusInternalServerError)
		return
	}

	activityID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to retrieve activity ID", http.StatusInternalServerError)
		return
	}

	for _, ActivityRole := range activity.ActivityRoles {
		_, err := tx.Exec("INSERT INTO ActivityRole (activity_id, role) VALUES (?, ?)", activityID, ActivityRole)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to insert activity role", http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
