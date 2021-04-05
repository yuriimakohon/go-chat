package service

import "errors"

var (
	authPrompt           = "auth-service: "
	roomPrompt           = "room-service: "
	ErrBadCredsFormat    = errors.New(authPrompt + "Invalid credentials format")
	ErrUserAlreadyExists = errors.New(authPrompt + "User already exists")
	ErrWrongPassword     = errors.New(authPrompt + "Wrong password")
	ErrUserDoesntExists  = errors.New(authPrompt + "User doesn't exists")
	ErrRoomAlreadyExists = errors.New(roomPrompt + "Room already exists")
	ErrRoomDoesntExists  = errors.New(roomPrompt + "Room doesn't exists")
	ErrUserAlreadyInRoom = errors.New(roomPrompt + "User already in room")
)
