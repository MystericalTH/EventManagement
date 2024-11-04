package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Feedback struct {
	feedbackID       int
	activityID       int
	memberID         int
	feedbackMessage  string
	feedbackDateTime time.Time
}

// Global variable for database connection

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
	// Query the database for feedback information
	var feedback Feedback
	err := db.QueryRow("SELECT feedbackID, activityID, memberID, feedbackMessage, feedbackDateTime FROM feedback WHERE feedbackID = ?", feedback.feedbackID).
		Scan(&feedback.feedbackID, &feedback.activityID, &feedback.memberID, &feedback.feedbackMessage, &feedback.feedbackDateTime)
	if err != nil {
		http.Error(w, "Failed to fetch feedback information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Output feedback info (this is a placeholder; adapt as needed)
	fmt.Fprintf(w, "Feedback: %+v\n", feedback)
}

func postFeedback(w http.ResponseWriter, r *http.Request) {
	// Parse form values
	activityID, err := strconv.Atoi(r.FormValue("activityID"))
	if err != nil {
		http.Error(w, "Invalid activityID", http.StatusBadRequest)
		return
	}
	memberID, err := strconv.Atoi(r.FormValue("memberID"))
	if err != nil {
		http.Error(w, "Invalid memberID", http.StatusBadRequest)
		return
	}
	feedbackMessage := r.FormValue("feedbackMessage")
	feedbackDateTime := time.Now()

	// Insert feedback information into the database
	_, err = db.Exec("INSERT INTO feedback (activityID, memberID, feedbackMessage, feedbackDateTime) VALUES (?, ?, ?, ?)", activityID, memberID, feedbackMessage, feedbackDateTime)
	if err != nil {
		http.Error(w, "Failed to insert feedback information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Feedback inserted successfully")
}

func feedback() {
	// Initialize the database connection (replace with your DSN)
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	http.HandleFunc("/feedback", feedbackHandler)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
