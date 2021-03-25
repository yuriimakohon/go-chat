package config

import "time"

const (
	BcryptCost  = 8
	TokenMaxAge = 20 * time.Second
	EnvPath     = "config/.env"
	HTMLPath    = "web/*.html"
	IndexPage   = "index.html"
	SignupPage  = "signup.html"
	LoginPage   = "login.html"
)
