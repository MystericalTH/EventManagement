package main

import (
	"fmt",
	"database/sql",
	"net/http",
	"html/template",
	"log",

	_ "github.com/go-sql-driver/mysql"
)

type Activity struct {
	activityID int(11),
	title varchar(255),
	startDate date,
	endDate date,
	maxNumber int(11),
	format varchar(10),
	description text,
	applicationStatus varchar(20),
	approvalStatus varchar(20),
	proposer int(11)
}

type Workshop struct {
	activityID int(11),
	startTime time,
	endTime time
}

type Project struct {
	activityID int(11),
	advisor varchar(255)
}

var (
	db *sql.DB
)

func getActivity(w http.ResponseWriter, r *http.Request) {
	// Query the database for activity information
	var activity Activity
	err := db.QueryRow("SELECT id, title, startDate, endDate, maxNumber, format, description, applicationStatus, 
			approvalStatus, proposer FROM activities WHERE id = ?", activity.activityID).Scan(&activity.activityID, &activity.title, &activity.startDate,
			&activity.endDate, &activity.maxNumber, &activity.format, &activity.description, &activity.applicationStatus, &activity.approvalStatus, &activity.proposer)
	if err != nil {
		http.Error(w, "Failed to fetch activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func postActivity(w http.ResponseWriter, r *http.Request) {
	// Get activity information from the request
	activity := Activity{
		activityID: r.FormValue("activityID"),
		title: r.FormValue("title"),
		startDate: r.FormValue("startDate"),
		endDate: r.FormValue("endDate"),
		maxNumber: r.FormValue("maxNumber"),
		format: r.FormValue("format"),
		description: r.FormValue("description"),
		applicationStatus: r.FormValue("applicationStatus"),
		approvalStatus: r.FormValue("approvalStatus"),
		proposer: r.FormValue("proposer"),
	}
	
	// Insert activity information into the database
	_, err := db.Exec("INSERT INTO activities (id, title, startDate, endDate, maxNumber, format, description, applicationStatus,
			approvalStatus, proposer) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", activity.activityID, activity.title, activity.startDate,
			activity.endDate, activity.maxNumber, activity.format, activity.description, activity.applicationStatus, activity.approvalStatus, activity.proposer)
	if err != nil {
		http.Error(w, "Failed to insert activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity page
	http.Redirect(w, r, "/activity", http.StatusFound)
	
}