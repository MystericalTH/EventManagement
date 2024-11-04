package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Define chatDevAd struct with JSON tags
type chatDevAd struct {
	MessageID   int       `json:"messageID"`
	AdminID     int       `json:"adminID"`
	DeveloperID int       `json:"developerID"`
	Message     string    `json:"message"`
	Datetime    time.Time `json:"datetime"`
}

// Define Developer struct with JSON tags
type Developer struct {
	DeveloperID int    `json:"developerID"`
	Email       string `json:"email"`
}

// Define Admin struct with JSON tags
type Admin struct {
	AdminID int    `json:"adminID"`
	Email   string `json:"email"`
}

func chat() {
	var err error
	// Initialize database connection
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/activitymanagement")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/chat", chatHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/developer", developerHandler)
	http.ListenAndServe(":8080", nil)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getChat(w, r)
	case "POST":
		postChat(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getChat(w http.ResponseWriter, r *http.Request) {
	var chat chatDevAd
	err := db.QueryRow("SELECT messageID, adminID, developerID, message, datetime FROM chat WHERE messageID = ?", r.URL.Query().Get("messageID")).Scan(
		&chat.MessageID, &chat.AdminID, &chat.DeveloperID, &chat.Message, &chat.Datetime)
	if err != nil {
		http.Error(w, "Failed to fetch chat information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}

func postChat(w http.ResponseWriter, r *http.Request) {
	var chat chatDevAd

	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	chat.Datetime = time.Now()

	_, err = db.Exec("INSERT INTO chat (adminID, developerID, message, datetime) VALUES (?, ?, ?, ?)", chat.AdminID, chat.DeveloperID, chat.Message, chat.Datetime)
	if err != nil {
		http.Error(w, "Failed to insert chat information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Chat message inserted successfully")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAdmin(w, r)
	case "POST":
		postAdmin(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAdmin(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	err := db.QueryRow("SELECT adminID, email FROM admin WHERE adminID = ?", r.URL.Query().Get("adminID")).Scan(&admin.AdminID, &admin.Email)
	if err != nil {
		http.Error(w, "Failed to fetch admin information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}

func postAdmin(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO admin (email) VALUES (?)", admin.Email)
	if err != nil {
		http.Error(w, "Failed to insert admin information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Admin inserted successfully")
}

func developerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDeveloper(w, r)
	case "POST":
		postDeveloper(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getDeveloper(w http.ResponseWriter, r *http.Request) {
	var developer Developer
	err := db.QueryRow("SELECT developerID, email FROM developer WHERE developerID = ?", r.URL.Query().Get("developerID")).Scan(&developer.DeveloperID, &developer.Email)
	if err != nil {
		http.Error(w, "Failed to fetch developer information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(developer)
}

func postDeveloper(w http.ResponseWriter, r *http.Request) {
	var developer Developer
	err := json.NewDecoder(r.Body).Decode(&developer)
	if err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO developer (email) VALUES (?)", developer.Email)
	if err != nil {
		http.Error(w, "Failed to insert developer information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Developer inserted successfully")
}
