package api

import (
	"sinno-server/pkg/api/handler"
	"sinno-server/pkg/db"

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
		}) // List members
		api.GET("/members/requests", func(c *gin.Context) {
			handler.GetAllMemberRequests(c, queries) // Pass queries to the handler
		}) // Accept a member
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
			handler.GetActivities(c, queries)
		})
		api.GET("/activities/:activityId", func(c *gin.Context) {
			handler.GetActivityByID(c, queries)
		})
		api.POST("/activities", func(c *gin.Context) {
			handler.PostActivity(c, queries)
		})
		api.GET("/activities/:activityId/roles", func(c *gin.Context) {
			handler.GetActivityRoles(c, queries)
		})

		// Feedback routes
		api.GET("/activities/:activityId/feedback/status", func(c *gin.Context) {
			handler.GetFeedbackStatus(c, queries)
		})
		api.POST("/activities/:activityId/feedback/submit", func(c *gin.Context) {
			handler.SubmitFeedback(c, queries)
		})

		// Registration routes
		api.GET("/activities/:activityId/registration/status", func(c *gin.Context) {
			handler.GetRegistrationStatus(c, queries)
		})
		api.POST("/activities/:activityId/registration/submit", func(c *gin.Context) {
			handler.SubmitRegistration(c, queries)
		})
		api.GET("/health", func(c *gin.Context) {
			handler.Healthchecks(c)
		})
	}
}
