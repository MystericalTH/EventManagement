package api

import (
	"net/http"
	api "sinno-server/pkg/api/handler"

	"github.com/gorilla/mux"
)

func LogRoutes() {
	http.HandleFunc("/api/login", HandleLogin)
	http.HandleFunc("/api/auth/google/callback", HandleCallback)
	http.HandleFunc("/api/logout", HandleLogout)
	http.HandleFunc("/api/verify", api.HandleVerifyRole)
}

func ActivitiesRoutes() {
	http.HandleFunc("/api/activities", api.GetActivities)
	http.HandleFunc("/api/activities/", api.GetActivityByID)
	http.HandleFunc("/api/proposal/submit", api.PostActivity)
}

func FeedbackRoutes(router *mux.Router) {
	router.HandleFunc("/api/activities/{activityId}/feedback/status", api.GetFeedbackStatus).Methods("GET")
}
