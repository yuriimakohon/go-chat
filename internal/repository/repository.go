package repository

import (
	"github.com/yuriimakohon/go-chat/internal/models"
)

type Authorization interface {
	CreateUser(creds models.Credentials) error
	GetUser(login string) (models.Credentials, error)
}

type Room interface {
	CreateRoom(room models.Room) error
	JoinRoom(login string, roomId string) error
}

type Message interface {
}

type Repository struct {
	Authorization
	Room
	Message
}

func NewRepository(auth Authorization, room Room) *Repository {
	return &Repository{
		Authorization: auth,
		Room:          room,
		Message:       nil,
	}
}
