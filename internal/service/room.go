package service

import (
	"github.com/yuriimakohon/go-chat/internal/models"
	"github.com/yuriimakohon/go-chat/internal/repository"
)

type RoomService struct {
	repo repository.Room
}

func NewRoomService(repo repository.Room) *RoomService {
	return &RoomService{repo: repo}
}

func (r *RoomService) CreateRoom(room models.Room) error {
	if err := r.repo.CreateRoom(room); err != nil {
		if err == repository.ErrRoomAlreadyExists {
			return ErrRoomAlreadyExists
		}
		return err
	}
	return nil
}

func (r *RoomService) JoinRoom(login string, roomId string) error {
	if err := r.repo.JoinRoom(login, roomId); err != nil {
		switch err {
		case repository.ErrRoomDosentExists:
			return ErrUserDosentExists
		case repository.ErrUserAlreadyInRoom:
			return ErrUserAlreadyInRoom
		default:
			return err
		}
	}
	return nil
}
