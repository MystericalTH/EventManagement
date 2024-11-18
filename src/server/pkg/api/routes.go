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
		api.GET("/members", func(c *gin.Context) {
			handler.GetAllMembers(c, queries) // Pass queries to the handler
		}) // List members
		api.GET("/members/:id", func(c *gin.Context) {
			handler.GetMemberByID(c, queries) // Pass queries to the handler
		}) // Get member by ID
		api.POST("/members", func(c *gin.Context) {
			handler.CreateMember(c, queries) // Pass queries to the handler
		}) // Create a new member
		api.PUT("/members/:id/accept", func(c *gin.Context) {
			handler.AcceptMember(c, queries) // Pass queries to the handler
		}) // Accept a member
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
	}
}
