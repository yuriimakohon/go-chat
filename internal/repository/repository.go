package repository

import (
	"github.com/yuriimakohon/go-chat/internal/models"
)

type Authorization interface {
	CreateUser(creds models.Credentials) error
	GetUser(login string) (models.Credentials, error)
}

type Room interface {
}

type Message interface {
}

type Repository struct {
	Authorization
	Room
	Message
}

func NewRepository(auth Authorization) *Repository {
	return &Repository{
		Authorization: auth,
		Room:          nil,
		Message:       nil,
	}
}
