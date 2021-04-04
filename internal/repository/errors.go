package repository

import "errors"

var (
	authPrompt           = "auth-repository: "
	roomPrompt           = "room-repository: "
	ErrUserAlreadyExists = errors.New(authPrompt + "User already exists")
	ErrUserDoesntExists  = errors.New(authPrompt + "User doesn't exists")
	ErrRoomAlreadyExists = errors.New(roomPrompt + "Room already exists")
	ErrRoomDosentExists  = errors.New(roomPrompt + "Room doesn't exists")
	ErrUserAlreadyInRoom = errors.New(roomPrompt + "User already in room")
)
