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
	router.HandleFunc("/api/activities/{id}", api.GetActivityByID).Methods("GET")
	router.HandleFunc("/api/activities", api.PostActivity).Methods("POST")
}

func FeedbackRoutes(router *mux.Router) {
	router.HandleFunc("/api/activities/{activityId}/feedback/status", api.GetFeedbackStatus).Methods("GET")
	router.HandleFunc("/api/feedback/submit", api.SubmitFeedback).Methods("POST")
}
