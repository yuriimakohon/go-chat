package configs

import "time"

const (
	TokenAge   = 3 * time.Minute
	EnvPath    = "configs/.env"
	HTMLPath   = "web/*.html"
	IndexPage  = "index.html"
	SignupPage = "signup.html"
	LoginPage  = "login.html"
)
