package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Define chatDevAd struct
type chatDevAd struct {
	messageID   int
	adminID     int
	developerID int
	message     string
	datetime    time.Time
}

// Define Developer struct
type Developer struct {
	developerID int
	email       string
}

// Define Admin struct
type Admin struct {
	adminID int
	email   string
}

func chat() {
	var err error
	// Initialize database connection
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/database_name")
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
		&chat.messageID, &chat.adminID, &chat.developerID, &chat.message, &chat.datetime)
	if err != nil {
		http.Error(w, "Failed to fetch chat information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Display chat information as an example
	fmt.Fprintf(w, "Chat message: %v", chat)
}

func postChat(w http.ResponseWriter, r *http.Request) {
	var chat chatDevAd

	// Convert FormValue strings to integers
	adminID, err := strconv.Atoi(r.FormValue("adminID"))
	if err != nil {
		http.Error(w, "Invalid adminID: "+err.Error(), http.StatusBadRequest)
		return
	}
	developerID, err := strconv.Atoi(r.FormValue("developerID"))
	if err != nil {
		http.Error(w, "Invalid developerID: "+err.Error(), http.StatusBadRequest)
		return
	}

	chat.adminID = adminID
	chat.developerID = developerID
	chat.message = r.FormValue("message")
	chat.datetime = time.Now()

	// Insert chat information into the database
	_, err = db.Exec("INSERT INTO chat (adminID, developerID, message, datetime) VALUES (?, ?, ?, ?)", chat.adminID, chat.developerID, chat.message, chat.datetime)
	if err != nil {
		http.Error(w, "Failed to insert chat information: "+err.Error(), http.StatusInternalServerError)
		return
	}

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
	err := db.QueryRow("SELECT adminID, email FROM admin WHERE adminID = ?", r.URL.Query().Get("adminID")).Scan(&admin.adminID, &admin.email)
	if err != nil {
		http.Error(w, "Failed to fetch admin information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Display admin information as an example
	fmt.Fprintf(w, "Admin email: %v", admin.email)
}

func postAdmin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	_, err := db.Exec("INSERT INTO admin (email) VALUES (?)", email)
	if err != nil {
		http.Error(w, "Failed to insert admin information: "+err.Error(), http.StatusInternalServerError)
		return
	}

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
	err := db.QueryRow("SELECT developerID, email FROM developer WHERE developerID = ?", r.URL.Query().Get("developerID")).Scan(&developer.developerID, &developer.email)
	if err != nil {
		http.Error(w, "Failed to fetch developer information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Display developer information as an example
	fmt.Fprintf(w, "Developer email: %v", developer.email)
}

func postDeveloper(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	_, err := db.Exec("INSERT INTO developer (email) VALUES (?)", email)
	if err != nil {
		http.Error(w, "Failed to insert developer information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Developer inserted successfully")
}
