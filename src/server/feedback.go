package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Feedback struct {
	FeedbackID       int       `json:"feedbackID"`
	ActivityID       int       `json:"activityID"`
	MemberID         int       `json:"memberID"`
	FeedbackMessage  string    `json:"feedbackMessage"`
	FeedbackDateTime time.Time `json:"feedbackDateTime"`
}

func feedback() {
	var err error
	// Initialize the database connection
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	http.HandleFunc("/feedback", feedbackHandler)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func feedbackHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getFeedback(w, r)
	case "POST":
		postFeedback(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback Feedback
	feedbackID := r.URL.Query().Get("feedbackID")

	// Validate feedbackID
	id, err := strconv.Atoi(feedbackID)
	if err != nil {
		http.Error(w, "Invalid feedback ID", http.StatusBadRequest)
		return
	}

	// Query the database for feedback information
	err = db.QueryRow("SELECT feedbackID, activityID, memberID, feedbackMessage, feedbackDateTime FROM feedback WHERE feedbackID = ?", id).
		Scan(&feedback.FeedbackID, &feedback.ActivityID, &feedback.MemberID, &feedback.FeedbackMessage, &feedback.FeedbackDateTime)
	if err != nil {
		http.Error(w, "Failed to fetch feedback information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feedback)
}

func postFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback Feedback

	// Decode JSON request body into feedback struct
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}
	feedback.FeedbackDateTime = time.Now() // Set current time for feedbackDateTime

	// Insert feedback information into the database
	_, err = db.Exec("INSERT INTO feedback (activityID, memberID, feedbackMessage, feedbackDateTime) VALUES (?, ?, ?, ?)",
		feedback.ActivityID, feedback.MemberID, feedback.FeedbackMessage, feedback.FeedbackDateTime)
	if err != nil {
		http.Error(w, "Failed to insert feedback information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Feedback inserted successfully")
}
