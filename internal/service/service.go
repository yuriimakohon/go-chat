package service

import (
	"github.com/yuriimakohon/go-chat/internal/models"
	"github.com/yuriimakohon/go-chat/internal/repository"
)

type Authorization interface {
	SignUp(creds models.Credentials) error
	LogIn(creds models.Credentials) error
}

type Room interface {
}

type Message interface {
}

type Service struct {
	Authorization
	Room
	Message
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Room:          nil,
		Message:       nil,
	}
}
