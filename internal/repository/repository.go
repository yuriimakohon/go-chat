package repository

import (
	"errors"
	"github.com/yuriimakohon/go-chat/internal/models/credentials"
)

var (
	ErrUserNotFound      = errors.New("repository: user not found")
	ErrUserAlreadyExists = errors.New("repository: user already exists")
)

type Repository interface {
	NewUser(cred credentials.Credentials) error
	GetUserByLogin(login string) (credentials.Credentials, error)
}
