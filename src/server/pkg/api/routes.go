package api

import (
	"sinno-server/pkg/api/handler"
	"sinno-server/pkg/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, queries *db.Queries) {

	api := router.Group("/api")
	{

		//! AUTH ROUTES !//

		api.GET("/login", func(c *gin.Context) {
			handler.AuthLogin(c)
		})
		api.GET("/auth/google/callback", func(c *gin.Context) {
			handler.AuthCallback(c, queries)
		})
		api.GET("/login/callback", func(c *gin.Context) {
			handler.LoginInfoRetrieval(c)
		})
		api.GET("/logout", func(c *gin.Context) {
			handler.AuthLogout(c)
		})

		//! MEMBER ROUTES !//

		// GET /members
		api.GET("/members", func(c *gin.Context) {
			handler.GetAllMembers(c, queries)
		})

		// GET /members/requests
		api.GET("/members/requests", func(c *gin.Context) {
			handler.GetAllMemberRequests(c, queries)
		})

		// DELETE /members/requests/:id delete a member request
		api.DELETE("/members/requests/:id", func(c *gin.Context) {
			handler.DeleteMember(c, queries)
		})

		// GET /members/:id delete a member
		api.DELETE("/members/:id", func(c *gin.Context) {
			handler.DeleteMember(c, queries)
		})

		// GET /members/:id get a member by id
		api.GET("/members/:id", func(c *gin.Context) {
			handler.GetMemberByID(c, queries)
		})

		// POST /members create a member
		api.POST("/members", func(c *gin.Context) {
			handler.CreateMember(c, queries)
		})

		// PUT /members/:id/accept accept a member request
		api.PUT("/members/:id/approve", func(c *gin.Context) {
			handler.AcceptMember(c, queries)
		})

		// PUT /members/:id update a member
		api.PUT("/members/:id", func(c *gin.Context) {
			handler.UpdateMember(c, queries)
		})

		//! ACTIVITY ROUTES !//

		// GET /activities get all requesting activities
		api.GET("/activities/requests", func(c *gin.Context) {
			handler.GetActivities(c, queries)
		})

		// GET /activities/:id get an activity by ID
		api.GET("/activities/:id", func(c *gin.Context) {
			handler.GetActivityByID(c, queries)
		})

		// POST /activities create an activity
		api.POST("/proposal/submit", func(c *gin.Context) {
			handler.PostActivity(c, queries)
		})

		// POST /activities/:id/roles create an activity role
		api.GET("/activities/:id/roles", func(c *gin.Context) {
			handler.GetActivityRoles(c, queries)
		})
		api.GET("/activities/", func(c *gin.Context) {
			handler.GetAcceptedActivities(c, queries)
		})

		//! FEEDBACK ROUTES !//

		// GET /feedback get a feedback
		api.GET("/activities/:id/feedback/status", func(c *gin.Context) {
			handler.GetFeedbackStatus(c, queries)
		})

		// POST /feedback submit a feedback
		api.POST("/activities/:id/feedback", func(c *gin.Context) {
			handler.SubmitFeedback(c, queries)
		})
		api.GET("/activities/:id/feedback", func(c *gin.Context) {
			handler.GetFeedbacksByActivity(c, queries)
		})

		// GET /registration status
		api.GET("/activities/:id/registration/status", func(c *gin.Context) {
			handler.GetRegistrationStatus(c, queries)
		})

		api.POST("/activities/:id/registration", func(c *gin.Context) {
			handler.SubmitRegistration(c, queries)
		})

		// Get /your proposed activities submitted response
		api.GET("/activities/:id/registration", func(c *gin.Context) {
			handler.GetSubmittedMembers(c, queries)
		})

		// 		api.GET("/activities/:id/registration", func(c *gin.Context) {
		// 			handler.GetActivityRegistration(c, queries)
		// 		})

		api.GET("/health", func(c *gin.Context) {
			handler.Healthchecks(c)
		})

		// Chat routes
		api.POST("/chats", func(c *gin.Context) {
			handler.CreateChat(c, queries)
		})
		api.GET("/chats", func(c *gin.Context) {
			handler.ListInitialAdminDevChat(c, queries)
		})
		api.GET("/chats/:id", func(c *gin.Context) {
			handler.ListAdminDevChats(c, queries)
		})

		// PUT /registration status
		api.PUT("/activities/:id/approve", func(c *gin.Context) {
			handler.ApproveActivityRegistration(c, queries)
		})

		// DELETE /registration
		api.DELETE("activities/:id", func(c *gin.Context) {
			handler.DeleteActivity(c, queries) // Pass queries to the handler
		})

		// GET /proposer's proposals

		api.GET("member/activities/proposals", func(c *gin.Context) {
			handler.GetProposerProposals(c, queries)
		})

		//! MEMBERS AND ACTIVITIES !//

		api.GET("member/activities", func(c *gin.Context) {
			handler.GetMemberActivities(c, queries)
		})

	}
}
