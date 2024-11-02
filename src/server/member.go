package main

import (
	"fmt",
	"database/sql",
	"net/http",
	"html/template",
	"log",

	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	memberID int(11),
	fName varchar(255),
	lName varchar(255),
	email varchar(320),
	phone varchar(20),
	githubUrl varchar(320),
	status varchar(10),
	reason text,
    acceptDateTime datetime,
    acceptAdmin int(11)
}

var (
	db *sql.DB
)

func memberHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getMember(w, r)
    case "POST":
        postMember(w, r)
    case "PATCH":
        patchMember(w, r)
    case "DELETE":
        deleteMember(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getMember(w http.ResponseWriter, r *http.Request) {
    // Query the database for member information
    var member Member
    err := db.QueryRow("SELECT id, fName, lName, email, phone, githubUrl, status, reason, acceptDateTime, 
        acceptAdmin FROM members WHERE id = ?", memberID).Scan(&member.memberID, &member.fName, &member.lName, &member.email, 
        &member.phone, &member.githubUrl, &member.status, &member.reason, &member.acceptDateTime, &member.acceptAdmin)
    if err != nil {
        http.Error(w, "Failed to fetch member information: "+err.Error(), http.StatusInternalServerError)
        return
    }
}

func postMember(w http.ResponseWriter, r *http.Request) {
    // Get form values
    fName := r.FormValue("fName")
    lName := r.FormValue("lName")
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    githubUrl := r.FormValue("githubUrl")
    status := r.FormValue("status")
    reason := r.FormValue("reason")
    acceptDateTime := r.FormValue("acceptDateTime")
    acceptAdmin := r.FormValue("acceptAdmin")

    // Insert member into database
    _, err = db.Exec("INSERT INTO members (fName, lName, email, phone, githubUrl, status, reason, acceptDateTime, acceptAdmin) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", fName, lName, email, phone, githubUrl, status, reason, acceptDateTime, acceptAdmin)
    if err != nil {
        http.Error(w, "Failed to insert member: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Redirect to profile
    http.Redirect(w, r, "/profile", http.StatusFound)
}

func patchMember(w http.ResponseWriter, r *http.Request) {
    // Get form values
    memberID := r.FormValue("memberID")
    fName := r.FormValue("fName")
    lName := r.FormValue("lName")
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    githubUrl := r.FormValue("githubUrl")
    status := r.FormValue("status")
    reason := r.FormValue("reason")

    // Build the SQL query dynamically based on provided fields
    query := "UPDATE members SET "
    params := []interface{}{}
    
    if fName != "" {
        query += "fName = ?, "
        params = append(params, fName)
    }
    if lName != "" {
        query += "lName = ?, "
        params = append(params, lName)
    }
    if email != "" {
        query += "email = ?, "
        params = append(params, email)
    }
    if phone != "" {
        query += "phone = ?, "
        params = append(params, phone)
    }
    if githubUrl != "" {
        query += "githubUrl = ?, "
        params = append(params, githubUrl)
    }
    if status != "" {
        query += "status = ?, "
        params = append(params, status)
    }
    if reason != "" {
        query += "reason = ?, "
        params = append(params, reason)
    }

    // Remove the trailing comma and space
    query = query[:len(query)-2]
    query += " WHERE id = ?"
    params = append(params, memberID)

    // Update member in database
    _, err := db.Exec(query, params...)
    if err != nil {
        http.Error(w, "Failed to update member: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Redirect to profile
    http.Redirect(w, r, "/profile", http.StatusFound)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
    // Get form values
    memberID := r.FormValue("memberID")

    // Delete member from database
    _, err = db.Exec("DELETE FROM members WHERE id = ?", memberID)
    if err != nil {
        http.Error(w, "Failed to delete member: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Redirect to profile
    http.Redirect(w, r, "/profile", http.StatusFound)
}