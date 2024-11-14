package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Activity struct {
	ActivityID        int       `json:"activityID"`
	Title             string    `json:"title"`
	Proposer          int       `json:"proposer"`
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	MaxNumber         int       `json:"maxNumber"`
	Format            string    `json:"format"`
	Description       string    `json:"description"`
	ProposeDateTime   time.Time `json:"proposeDateTime"`
	AcceptAdmin       int       `json:"acceptAdmin"`
	AcceptDateTime    time.Time `json:"acceptDateTime"`
	ApplicationStatus string    `json:"applicationStatus"`
}

func activity() {
	// Initialize the database connection
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/activitymanagement")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Setup routes
	http.HandleFunc("/activity", activityHandler)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getActivity(w, r)
	case "POST":
		postActivity(w, r)
	case "PATCH":
		patchActivity(w, r)
	case "DELETE":
		deleteActivity(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getActivity(w http.ResponseWriter, r *http.Request) {
	activityID := r.URL.Query().Get("activityID")
	if activityID == "" {
		http.Error(w, "activityID not provided", http.StatusBadRequest)
		return
	}

	var activity Activity
	err := db.QueryRow("SELECT id, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus FROM activities WHERE id = ?", activityID).
		Scan(&activity.ActivityID, &activity.Title, &activity.Proposer, &activity.StartDate, &activity.EndDate, &activity.MaxNumber, &activity.Format, &activity.Description, &activity.ProposeDateTime, &activity.AcceptAdmin, &activity.AcceptDateTime, &activity.ApplicationStatus)
	if err != nil {
		http.Error(w, "Failed to fetch activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
}

func postActivity(w http.ResponseWriter, r *http.Request) {
	var activity Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO activities (id, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		activity.ActivityID, activity.Title, activity.Proposer, activity.StartDate, activity.EndDate, activity.MaxNumber, activity.Format, activity.Description, activity.ProposeDateTime, activity.AcceptAdmin, activity.AcceptDateTime, activity.ApplicationStatus)
	if err != nil {
		http.Error(w, "Failed to insert activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func patchActivity(w http.ResponseWriter, r *http.Request) {
	var activity Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE activities SET title = ?, proposer = ?, startDate = ?, endDate = ?, maxNumber = ?, format = ?, description = ?, proposeDateTime = ?, acceptAdmin = ?, acceptDateTime = ?, applicationStatus = ? WHERE id = ?",
		activity.Title, activity.Proposer, activity.StartDate, activity.EndDate, activity.MaxNumber, activity.Format, activity.Description, activity.ProposeDateTime, activity.AcceptAdmin, activity.AcceptDateTime, activity.ApplicationStatus, activity.ActivityID)
	if err != nil {
		http.Error(w, "Failed to update activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteActivity(w http.ResponseWriter, r *http.Request) {
	activityID, err := strconv.Atoi(r.URL.Query().Get("activityID"))
	if err != nil {
		http.Error(w, "Invalid activityID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM activities WHERE id = ?", activityID)
	if err != nil {
		http.Error(w, "Failed to delete activity: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
