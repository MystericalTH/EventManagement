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
	proposer int(11),
	startDate date,
	endDate date,
	maxNumber int(11),
	format varchar(10),
	description text,
	proposeDateTime datetime,
	acceptAdmin int(11),
	acceptDateTime datetime,
	applicationStatus varchar(20)
}

type Workshop struct {
	workshopID int(11),
	startTime time,
	endTime time
}

type Project struct {
	projectID int(11),
	advisor varchar(255)
}

var (
	db *sql.DB
)

func activityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getActivity(w, r)
	case "POST":
		postActivity(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getActivity(w http.ResponseWriter, r *http.Request) {
	// Query the database for activity information
	var activity Activity
	err := db.QueryRow("SELECT id, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, 
			acceptAdmin, acceptDateTime, applicationStatus FROM activities WHERE id = ?", activity.activityID).Scan(&activity.activityID, &activity.title, &activity.proposer, &activity.startDate,
			&activity.endDate, &activity.maxNumber, &activity.format, &activity.description, &activity.proposeDateTime, &activity.acceptAdmin, &activity.acceptDateTime, &activity.applicationStatus)
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
		proposer: r.FormValue("proposer"),
		startDate: r.FormValue("startDate"),
		endDate: r.FormValue("endDate"),
		maxNumber: r.FormValue("maxNumber"),
		format: r.FormValue("format"),
		description: r.FormValue("description"),
		proposeDateTime: r.FormValue("proposeDateTime"),
		acceptAdmin: r.FormValue("acceptAdmin"),
		acceptDateTime: r.FormValue("acceptDateTime"),
		applicationStatus: r.FormValue("applicationStatus")
	}
	
	// Insert activity information into the database
	_, err := db.Exec("INSERT INTO activities (id, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime,
			acceptAdmin, acceptDateTime, applicationStatus) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", activity.activityID, activity.title, activity.proposer, 
			activity.startDate, activity.endDate, activity.maxNumber, activity.format, activity.description, activity.proposeDateTime, activity.acceptAdmin, activity.acceptDateTime, activity.applicationStatus)
	if err != nil {
		http.Error(w, "Failed to insert activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity page
	http.Redirect(w, r, "/activity", http.StatusFound)
}