package repository

import "errors"

var (
	authPrompt           = "auth-repository: "
	ErrUserAlreadyExists = errors.New(authPrompt + "User already exists")
	ErrUserDoesntExists  = errors.New(authPrompt + "User doesn't exists")
)
