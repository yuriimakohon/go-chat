package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	UserLogin string    `json:"user_login"`
	RoomId    string    `json:"room_id"`
	Id        uuid.UUID `json:"id"`
	Time      time.Time `json:"time"`
}
