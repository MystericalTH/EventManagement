package api

import (
	"net/http"
)

func ActivitiesRoutes() {
	http.HandleFunc("/api/activities", GetActivities)
}

func LogRoutes() {
	http.HandleFunc("/api/login", HandleLogin)
	http.HandleFunc("/api/auth/google/callback", HandleCallback)
	http.HandleFunc("/api/logout", HandleLogout)
	http.HandleFunc("/api/verify", HandleVerifyRole)
}
