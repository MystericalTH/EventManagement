package handler

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
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

// AuthLogin redirects the user to Google OAuth consent page
func AuthLogin(c *gin.Context) {
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

// AuthCallback is called after OAuth2 authentication with Google
func AuthCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.String(http.StatusBadRequest, "Authorization code not found")
		return
	}

	// Exchange code for access token
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange token: %s", err.Error())
		return
	}

	// Retrieve user info from Google API
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
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		c.String(http.StatusInternalServerError, "Failed to decode user info: %s", err.Error())
		return
	}

	// Save user info in session
	session, err := sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve session: %s", err.Error())
		return
	}

	session.Values["user_email"] = userInfo.Email
	session.Values["user_name"] = userInfo.Name
	session.Save(c.Request, c.Writer)

	// Redirect to the specified or default URI
	redirectUri := c.DefaultQuery("redirect_uri", "/")
	c.Redirect(http.StatusFound, redirectUri)
}

// LoginInfoRetrieval retrieves login info from the session
func LoginInfoRetrieval(c *gin.Context) {
	session, err := sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication failed",
			"user":    nil,
			"role":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful",
		"user":    session.Values["user_name"],
		"role":    session.Values["role"], // You may want to store and retrieve this as needed
	})
}

// AuthLogout clears the session and redirects to home
func AuthLogout(c *gin.Context) {
	session, err := sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to retrieve session: %s", err.Error())
		return
	}
	session.Options.MaxAge = -1 // Clear the session
	session.Save(c.Request, c.Writer)

	// Redirect to home page
	c.Redirect(http.StatusFound, "/")
}

// HandleVerifyRole checks the user's role in the session
func HandleVerifyRole(c *gin.Context) {
	session, err := sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"role": "unknown"})
		return
	}

	role, ok := session.Values["role"].(string)
	if !ok {
		role = "unknown"
	}
	c.JSON(http.StatusOK, gin.H{"role": role})
}
