package api

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

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

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect user to Google's OAuth consent page
	role := r.URL.Query().Get("role")
	state := uuid.NewString() + "__" + role
	stateUrl := oauthConfig.AuthCodeURL(url.QueryEscape(state), oauth2.AccessTypeOffline)
	http.Redirect(w, r, stateUrl, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	// Retrieve the authorization code from Google
	code := r.URL.Query().Get("code")
	state, err := url.QueryUnescape(r.URL.Query().Get("state"))
	if err != nil {
		http.Error(w, "Failed to unescape state: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if code == "" {
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		return
	}

	// Exchange the code for a token
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the token to fetch user info from Google
	client := oauthConfig.Client(context.Background(), token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer userInfoResp.Body.Close()

	// Parse the user info
	var userInfo UserInfo
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	stateparts := strings.SplitN(state, "__", 2)
	role := stateparts[1]
	// Store user info in session
	session, _ := sessionStore.Get(r, sessionName)
	session.Values["user"] = userInfo
	session.Values["role"] = role

	session.Save(r, w)

	// Redirect to profile
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	// Clear the session
	session, _ := sessionStore.Get(r, sessionName)
	session.Options.MaxAge = -1
	session.Save(r, w)

	// Redirect to home
	http.Redirect(w, r, "/", http.StatusFound)
}
