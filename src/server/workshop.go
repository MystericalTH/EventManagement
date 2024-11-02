package main

import (
	"fmt",
	"database/sql",
	"net/http",
	"html/template",
	"log",

	_ "github.com/go-sql-driver/mysql"
)

type Workshop struct {
	workshopID int(11),
	startTime time,
	endTime time
}

var (
	db *sql.DB
)

func workshopHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getWorkshop(w, r)
	case "POST":
		postWorkshop(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getWorkshop(w http.ResponseWriter, r *http.Request) {
	// Query the database for workshop information
	var workshop Workshop
	err := db.QueryRow("SELECT id, startTime, endTime FROM workshop WHERE id = ?", workshop.workshopID).Scan(&workshop.workshopID, &workshop.startTime, &workshop.endTime)
	if err != nil {
		http.Error(w, "Failed to fetch workshop information: "+err.Error(), http.StatusInternalServerError)
		return
}

func postWorkshop(w http.ResponseWriter, r *http.Request) {
	// Insert workshop information into the database
	var workshop Workshop
	err := db.QueryRow("INSERT INTO workshop (startTime, endTime) VALUES (?, ?)", workshop.startTime, workshop.endTime)
	if err != nil {
		http.Error(w, "Failed to insert workshop information: "+err.Error(), http.StatusInternalServerError)
		return
	}
}