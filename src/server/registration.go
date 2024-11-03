package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type pjRegist struct {
	projectID int(11),
	memberID int(11),
	role varchar(30),
	expectation text,
	datetime datetime
}

type wsRegist struct {
	workshopID int(11),
	memberID int(11),
	role varchar(30),
	expectation text,
	datetime datetime
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPjRegist(w, r)
		getWsRegist(w, r)
	case "POST":
		postPjRegist(w, r)
		postWsRegist(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getPjRegist(w http.ResponseWriter, r *http.Request) {
	// Query the database for project registration information
	var pjRegist pjRegist
	err := db.QueryRow("SELECT projectID, memberID, role, expectation, datetime FROM projectRegistration WHERE projectID = ?", projectID).Scan(&pjRegist.projectID, &pjRegist.memberID, &pjRegist.role, &pjRegist.expectation, &pjRegist.datetime)
	if err != nil {
		http.Error(w, "Failed to fetch project registration information: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func getWsRegist(w http.ResponseWriter, r *http.Request) {
	// Query the database for workshop registration information
	var wsRegist wsRegist
	err := db.QueryRow("SELECT workshopID, memberID, role, expectation, datetime FROM workshopRegistration WHERE workshopID = ?", workshopID).Scan(&wsRegist.workshopID, &wsRegist.memberID, &wsRegist.role, &wsRegist.expectation, &wsRegist.datetime)
	if err != nil {
		http.Error(w, "Failed to fetch workshop registration information: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func postPjRegist(w http.ResponseWriter, r *http.Request) {
	// Get project registration information from the request
	pjRegist := pjRegist{
		projectID: r.FormValue("projectID"),
		memberID: r.FormValue("memberID"),
		role: r.FormValue("role"),
		expectation: r.FormValue("expectation"),
		datetime: r.FormValue("datetime")
	}
}

func postWsRegist(w http.ResponseWriter, r *http.Request) {
	// Get workshop registration information from the request
	wsRegist := wsRegist{
		projectID: r.FormValue("projectID"),
		memberID: r.FormValue("memberID"),
		role: r.FormValue("role"),
		expectation: r.FormValue("expectation"),
		datetime: r.FormValue("datetime")
	}
}