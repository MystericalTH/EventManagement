package api

import (
	"net/http"
	api "sinno-server/pkg/api/handler"
)

func ActivitiesRoutes() {
	http.HandleFunc("/api/activities", api.GetActivities)
}

func LogRoutes() {
	http.HandleFunc("/api/login", HandleLogin)
	http.HandleFunc("/api/auth/google/callback", HandleCallback)
	http.HandleFunc("/api/logout", HandleLogout)
	http.HandleFunc("/api/verify", api.HandleVerifyRole)
}
