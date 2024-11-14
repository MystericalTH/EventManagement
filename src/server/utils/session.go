package utils

import (
	"os"

	"github.com/gorilla/sessions"
)

var SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
