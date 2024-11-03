package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Feedback struct {
	feedbackID int(11),
	activityID int(11),
	memberID int(11),
	feedbackMessage text,
	feedbackDateTime datetime
}

var (
	db *sql.DB
)

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
	err := db.QueryRow("SELECT id, activityID, memberID, feedbackMessage, feedbackDateTime FROM feedback WHERE id = ?", feedback.feedbackID).Scan(&feedback.feedbackID, &feedback.activityID, &feedback.memberID, &feedback.feedbackMessage, &feedback.feedbackDateTime)
	if err != nil {
		http.Error(w, "Failed to fetch feedback information: "+err.Error(), http.StatusInternalServerError)
		return
}

func postFeedback(w http.ResponseWriter, r *http.Request) {
	// Insert feedback information into the database
	var feedback Feedback
	err := db.QueryRow("INSERT INTO feedback (activityID, memberID, feedbackMessage, feedbackDateTime) VALUES (?, ?, ?, ?)", feedback.activityID, feedback.memberID, feedback.feedbackMessage, feedback.feedbackDateTime)
	if err != nil {
		http.Error(w, "Failed to insert feedback information: "+err.Error(), http.StatusInternalServerError)
		return
	}
}