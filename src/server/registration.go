package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Project Registration struct
type PjRegist struct {
	ProjectID   int
	MemberID    int
	Role        string
	Expectation string
	Datetime    string
}

// Workshop Registration struct
type WsRegist struct {
	WorkshopID  int
	MemberID    int
	Role        string
	Expectation string
	Datetime    string
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
	projectID := r.URL.Query().Get("projectID")
	if projectID == "" {
		http.Error(w, "projectID not provided", http.StatusBadRequest)
		return
	}

	var pjRegist PjRegist
	err := db.QueryRow("SELECT projectID, memberID, role, expectation, datetime FROM projectRegistration WHERE projectID = ?", projectID).
		Scan(&pjRegist.ProjectID, &pjRegist.MemberID, &pjRegist.Role, &pjRegist.Expectation, &pjRegist.Datetime)
	if err != nil {
		http.Error(w, "Failed to fetch project registration information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Output fetched data, e.g., in JSON format or render with a template
	fmt.Fprintf(w, "Project Registration: %+v", pjRegist)
}

func getWsRegist(w http.ResponseWriter, r *http.Request) {
	workshopID := r.URL.Query().Get("workshopID")
	if workshopID == "" {
		http.Error(w, "workshopID not provided", http.StatusBadRequest)
		return
	}

	var wsRegist WsRegist
	err := db.QueryRow("SELECT workshopID, memberID, role, expectation, datetime FROM workshopRegistration WHERE workshopID = ?", workshopID).
		Scan(&wsRegist.WorkshopID, &wsRegist.MemberID, &wsRegist.Role, &wsRegist.Expectation, &wsRegist.Datetime)
	if err != nil {
		http.Error(w, "Failed to fetch workshop registration information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Output fetched data, e.g., in JSON format or render with a template
	fmt.Fprintf(w, "Workshop Registration: %+v", wsRegist)
}

func postPjRegist(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	projectID, err := strconv.Atoi(r.FormValue("projectID"))
	if err != nil {
		http.Error(w, "Invalid projectID", http.StatusBadRequest)
		return
	}

	memberID, err := strconv.Atoi(r.FormValue("memberID"))
	if err != nil {
		http.Error(w, "Invalid memberID", http.StatusBadRequest)
		return
	}

	role := r.FormValue("role")
	expectation := r.FormValue("expectation")
	datetime := r.FormValue("datetime")

	_, err = db.Exec("INSERT INTO projectRegistration (projectID, memberID, role, expectation, datetime) VALUES (?, ?, ?, ?, ?)",
		projectID, memberID, role, expectation, datetime)
	if err != nil {
		http.Error(w, "Failed to insert project registration: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Project registration created successfully"))
}

func postWsRegist(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	workshopID, err := strconv.Atoi(r.FormValue("workshopID"))
	if err != nil {
		http.Error(w, "Invalid workshopID", http.StatusBadRequest)
		return
	}

	memberID, err := strconv.Atoi(r.FormValue("memberID"))
	if err != nil {
		http.Error(w, "Invalid memberID", http.StatusBadRequest)
		return
	}

	role := r.FormValue("role")
	expectation := r.FormValue("expectation")
	datetime := r.FormValue("datetime")

	_, err = db.Exec("INSERT INTO workshopRegistration (workshopID, memberID, role, expectation, datetime) VALUES (?, ?, ?, ?, ?)",
		workshopID, memberID, role, expectation, datetime)
	if err != nil {
		http.Error(w, "Failed to insert workshop registration: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Workshop registration created successfully"))
}

func registration() {
	// Initialize the database connection (replace with your DSN)
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	http.HandleFunc("/registration", registrationHandler)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
