package configs

import "time"

const (
	TokenMaxAge = 20 * time.Second
	EnvPath     = "configs/.env"
	HTMLPath    = "web/*.html"
	IndexPage   = "index.html"
	SignupPage  = "signup.html"
	LoginPage   = "login.html"
)
