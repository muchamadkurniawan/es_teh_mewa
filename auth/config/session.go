package config

import (
	"github.com/gorilla/sessions"
	"os"
)

const SESSION_ID = "go_auth_sess"

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
