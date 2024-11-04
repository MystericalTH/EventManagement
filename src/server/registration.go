package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Project Registration struct
type PjRegist struct {
	ProjectID   int    `json:"projectID"`
	MemberID    int    `json:"memberID"`
	Role        string `json:"role"`
	Expectation string `json:"expectation"`
	Datetime    string `json:"datetime"`
}

// Workshop Registration struct
type WsRegist struct {
	WorkshopID  int    `json:"workshopID"`
	MemberID    int    `json:"memberID"`
	Role        string `json:"role"`
	Expectation string `json:"expectation"`
	Datetime    string `json:"datetime"`
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

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPjRegist(w, r)
	case "POST":
		postPjRegist(w, r)
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

	// Return the data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pjRegist)
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

	// Return the data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wsRegist)
}

func postPjRegist(w http.ResponseWriter, r *http.Request) {
	var pjRegist PjRegist
	if err := json.NewDecoder(r.Body).Decode(&pjRegist); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO projectRegistration (projectID, memberID, role, expectation, datetime) VALUES (?, ?, ?, ?, ?)",
		pjRegist.ProjectID, pjRegist.MemberID, pjRegist.Role, pjRegist.Expectation, pjRegist.Datetime)
	if err != nil {
		http.Error(w, "Failed to insert project registration: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Project registration created successfully"))
}

func postWsRegist(w http.ResponseWriter, r *http.Request) {
	var wsRegist WsRegist
	if err := json.NewDecoder(r.Body).Decode(&wsRegist); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO workshopRegistration (workshopID, memberID, role, expectation, datetime) VALUES (?, ?, ?, ?, ?)",
		wsRegist.WorkshopID, wsRegist.MemberID, wsRegist.Role, wsRegist.Expectation, wsRegist.Datetime)
	if err != nil {
		http.Error(w, "Failed to insert workshop registration: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Workshop registration created successfully"))
}
