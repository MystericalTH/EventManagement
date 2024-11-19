package handler

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// Replace with your actual Google credentials
	port         = os.Getenv("LISTEN_PORT")
	redirectPort = os.Getenv("REDIRECT_PORT")
	oauthConfig  = &oauth2.Config{
		RedirectURL:  fmt.Sprintf("http://localhost:%s/api/auth/google/callback", redirectPort),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),     // Set in your environment
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"), // Set in your environment
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	sessionName  = "session-one"
)

type UserInfo struct {
	Name  string
	Email string
}

func init() {
	// Register UserInfo type to store in session
	gob.Register(UserInfo{})
}

func AuthLogin(c *gin.Context) {
	// Redirect user to Google's OAuth consent page
	role := c.Query("role")
	redirectUri, err := url.QueryUnescape(c.Query("redirect_uri"))
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to unescape redirectUri: %s", err.Error())
		return
	}

	state := uuid.NewString() + "__" + role + "__" + redirectUri
	stateUrl := oauthConfig.AuthCodeURL(url.QueryEscape(state), oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, stateUrl)
}

func AuthCallback(c *gin.Context) {
	// Retrieve authorization code and state
	code := c.Query("code")
	if code == "" {
		c.String(http.StatusBadRequest, "Authorization code not found")
		return
	}

	// Exchange the authorization code for an access token
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange token: %s", err.Error())
		return
	}

	// Retrieve user information
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create request: %s", err.Error())
		return
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := oauthConfig.Client(context.Background(), token)
	userInfoResp, err := client.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get user info: %s", err.Error())
		return
	}
	defer userInfoResp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		c.String(http.StatusInternalServerError, "Failed to decode user info: %s", err.Error())
		return
	}

	// Save user email in session
	session, _ := sessionStore.Get(c.Request, sessionName)
	session.Values["user_email"] = userInfo.Email
	if err := session.Save(c.Request, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to save session: %s", err.Error())
		return
	}

	// Redirect back
	redirectUri, _ := c.GetQuery("redirect_uri")
	if redirectUri == "" {
		redirectUri = "/"
	}
	c.Redirect(http.StatusFound, redirectUri)
}

func LoginInfoRetrieval(c *gin.Context) {
	session, err := sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication failed",
			"user":    nil,
			"role":    nil,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful",
		"user":    session.Values["user"],
		"role":    session.Values["role"],
	})
}

func AuthLogout(c *gin.Context) {
	// Clear the session
	session, _ := sessionStore.Get(c.Request, sessionName)
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	// Redirect to home
	c.Redirect(http.StatusFound, "/")
}

func HandleVerifyRole(c *gin.Context) {
	session, _ := sessionStore.Get(c.Request, sessionName)
	role, ok := session.Values["role"].(string)
	if !ok {
		role = "unknown"
	}
	c.JSON(http.StatusOK, gin.H{"role": role})
}
