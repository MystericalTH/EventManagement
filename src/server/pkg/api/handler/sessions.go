package handler

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"sinno-server/pkg/db"

	"log"

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
	SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	SessionName  = "session-one"
)

type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func init() {
	// Register UserInfo type to store in session
	gob.Register(UserInfo{})
}

// AuthLogin redirects the user to Google OAuth consent page
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

// AuthCallback is called after OAuth2 authentication with Google
func AuthCallback(c *gin.Context, queries *db.Queries) {
	// Retrieve the authorization code from Google
	code := c.Query("code")
	state, err := url.QueryUnescape(c.Query("state"))
	if err != nil {
		log.Printf("Error unescaping state: %v", err)
		c.String(http.StatusInternalServerError, "Failed to unescape state: %s", err.Error())
		return
	}

	if code == "" {
		log.Println("Authorization code not found")
		c.String(http.StatusBadRequest, "Authorization code not found")
		return
	}

	log.Printf("Authorization code received: %s, State: %s", code, state)

	// Exchange the code for a token
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging token: %v", err)
		c.String(http.StatusInternalServerError, "Failed to exchange token: %s", err.Error())
		return
	}

	log.Printf("Token successfully exchanged: %s", token.AccessToken)

	// Fetch user info
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		log.Printf("Error creating request for user info: %v", err)
		c.String(http.StatusInternalServerError, "Failed to create request: %s", err.Error())
		return
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := oauthConfig.Client(context.Background(), token)
	userInfoResp, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching user info: %v", err)
		c.String(http.StatusInternalServerError, "Failed to get user info: %s", err.Error())
		return
	}
	defer userInfoResp.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		log.Printf("Error decoding user info: %v", err)
		c.String(http.StatusInternalServerError, "Failed to decode user info: %s", err.Error())
		return
	}

	log.Printf("User info retrieved: %+v", userInfo)

	// Split the state to get role and redirectUri
	stateparts := strings.SplitN(state, "__", 3)
	if len(stateparts) < 3 {
		log.Printf("Invalid state format: %s", state)
		c.String(http.StatusBadRequest, "Invalid state parameter")
		return
	}

	role := stateparts[1]
	redirectUri := stateparts[2]
	log.Printf("Role: %s, Redirect URI: %s", role, redirectUri)

	// Role-based checks
	switch role {
	case "member":
		memberID, err := queries.GetMemberIDByEmail(context.Background(), userInfo.Email)
		if err != nil {
			log.Printf("Member not found: %s", userInfo.Email)
			c.String(http.StatusUnauthorized, "User not found for role: member, please go to signup page")
			return
		}

		_, err = queries.GetMemberIDByEmailWaitingAccept(context.Background(), userInfo.Email)
		if err != nil {
			log.Printf("Member waiting for acceptance: %s", userInfo.Email)
			c.String(http.StatusUnauthorized, "Please wait until admin accepts your request")
			return
		}

		log.Printf("Member validated with ID: %d", memberID)

	case "admin":
		if _, err := queries.GetAdminIDByEmail(context.Background(), userInfo.Email); err != nil {
			log.Printf("Admin not found: %s", userInfo.Email)
			c.String(http.StatusUnauthorized, "User not found for role: admin")
			return
		}
		log.Printf("Admin validated: %s", userInfo.Email)

	case "developer":
		if _, err := queries.GetDeveloperIDByEmail(context.Background(), userInfo.Email); err != nil {
			log.Printf("Developer not found: %s", userInfo.Email)
			c.String(http.StatusUnauthorized, "User not found for role: developer")
			return
		}
		log.Printf("Developer validated: %s", userInfo.Email)

	default:
		log.Printf("Invalid role: %s", role)

	}

	// Save user info in session
	session, err := SessionStore.Get(c.Request, SessionName)
	if err != nil {
		log.Printf("Error retrieving session: %v", err)
		c.String(http.StatusInternalServerError, "Failed to retrieve session: %s", err.Error())
		return
	}

	session.Values["user"] = userInfo
	session.Values["role"] = role
	session.Values["user_email"] = userInfo.Email
	session.Values["user_name"] = userInfo.Name
	if err := session.Save(c.Request, c.Writer); err != nil {
		log.Printf("Error saving session: %v", err)
		c.String(http.StatusInternalServerError, "Failed to save session: %s", err.Error())
		return
	}

	log.Printf("Session saved successfully for user: %s", userInfo.Email)

	// Redirect to the specified URI
	log.Printf("Redirecting to: %s", redirectUri)
	c.Redirect(http.StatusFound, redirectUri)
}

// LoginInfoRetrieval retrieves login info from the session
func LoginInfoRetrieval(c *gin.Context) {
	session, err := SessionStore.Get(c.Request, SessionName)
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
		"user":    session.Values["user"],
		"role":    session.Values["role"], // You may want to store and retrieve this as needed
	})
}

// AuthLogout clears the session and redirects to home
func AuthLogout(c *gin.Context) {
	session, err := SessionStore.Get(c.Request, SessionName)
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
	session, err := SessionStore.Get(c.Request, SessionName)
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

func Healthchecks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"healthchecks": "running"})
}
