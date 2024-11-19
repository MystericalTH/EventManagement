package api

import (
	"sinno-server/pkg/api/handler"
	"sinno-server/pkg/db"

	"github.com/gorilla/mux"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, queries *db.Queries) {
	// Comment out middleware temporarily for testing
	// router.Use(middleware.AuthMiddleware)

	// Define API groups or routes
	api := router.Group("/api")
	{
		// Member routes
		api.GET("/members", func(c *gin.Context) {
			handler.GetAllMembers(c, queries) // Pass queries to the handler
		})
		api.GET("/members/:id", func(c *gin.Context) {
			handler.GetMemberByID(c, queries) // Pass queries to the handler
		})
		api.POST("/members", func(c *gin.Context) {
			handler.CreateMember(c, queries) // Pass queries to the handler
		})
		api.PUT("/members/:id/accept", func(c *gin.Context) {
			handler.AcceptMember(c, queries) // Pass queries to the handler
		})

		// Authentication routes
		api.GET("/login", func(c *gin.Context) {
			handler.AuthLogin(c)
		})
		api.GET("/auth/google/callback", func(c *gin.Context) {
			handler.AuthCallback(c)
		})
		api.GET("/login/callback", func(c *gin.Context) {
			handler.LoginInfoRetrieval(c)
		})
		api.GET("/logout", func(c *gin.Context) {
			handler.AuthLogout(c)
		})

		// Activity routes
		api.GET("/activities", func(c *gin.Context) {
			handler.GetActivities(c)
		})
		api.GET("/activities/:id", func(c *gin.Context) {
			handler.GetActivityByID(c)
		})
		api.POST("/activities", func(c *gin.Context) {
			handler.PostActivity(c)
		})
		api.GET("/activities/:id/roles", func(c *gin.Context) {
			handler.GetActivityRoles(c)
		})
	}
}

func FeedbackRoutes(router *mux.Router) {
	router.HandleFunc("/api/activities/{activityId}/feedback/status", handler.GetFeedbackStatus).Methods("GET")
	router.HandleFunc("/api/activities/{activityId}/feedback/submit", handler.SubmitFeedback).Methods("POST")
}

func RegistrationRoutes(router *mux.Router) {
	router.HandleFunc("/api/activities/{activityId}/registration/status", handler.GetRegistrationStatus).Methods("GET")
	router.HandleFunc("/api/activities/{activityId}/registration/submit", handler.SubmitRegistration).Methods("POST")
}
