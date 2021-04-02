package service

import "errors"

var (
	authPrompt           = "auth-service: "
	ErrUserAlreadyExists = errors.New(authPrompt + "User already exists")
	ErrWrongPassword     = errors.New(authPrompt + "Wrong password")
	ErrUserDosentExists  = errors.New(authPrompt + "User doesn't exists")
)
