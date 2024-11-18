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

func ActivitiesRoutes(router *mux.Router) {
	router.HandleFunc("/api/activities", api.GetActivities).Methods("GET")
	router.HandleFunc("/api/activities/{activityId}", api.GetActivityByID).Methods("GET")
	router.HandleFunc("/api/activities", api.PostActivity).Methods("POST")
	router.HandleFunc("/api/activities/{activityId}/roles", api.GetActivityRoles).Methods("GET")
}

func FeedbackRoutes(router *mux.Router) {
	router.HandleFunc("/api/activities/{activityId}/feedback/status", api.GetFeedbackStatus).Methods("GET")
	router.HandleFunc("/api/activities/{activityId}/feedback/submit", api.SubmitFeedback).Methods("POST")
}

func RegistrationRoutes(router *mux.Router) {
	router.HandleFunc("/api/activities/{activityId}/registration/status", api.GetRegistrationStatus).Methods("GET")
	router.HandleFunc("/api/activities/{activityId}/registration/submit", api.SubmitRegistration).Methods("POST")
}
