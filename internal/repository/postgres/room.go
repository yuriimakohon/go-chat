package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/yuriimakohon/go-chat/internal/models"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"log"
)

type RoomRepository struct {
	db *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r RoomRepository) CreateRoom(room models.Room) error {
	if exists, err := r.isRoomExists(room.Id); exists {
		return repository.ErrRoomAlreadyExists
	} else if err != nil {
		return err
	}

	_, err := r.db.Exec("INSERT INTO rooms VALUES ($1)", room.Id)
	if err != nil {
		log.Printf("postgres/room - CreateRoom: %s\n", err)
		return err
	}
	return nil
}

func (r RoomRepository) JoinRoom(login string, roomId string) error {
	if exists, err := r.isRoomExists(roomId); err != nil {
		return err
	} else if !exists {
		return repository.ErrRoomDoesntExists
	}

	err := r.db.QueryRow(
		"SELECT FROM members WHERE user_login=$1 AND room_id=$2",
		login, roomId).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("postgres/room - JoinRoom: %v\n", err)
			return err
		}
	} else {
		return repository.ErrUserAlreadyInRoom
	}

	_, err = r.db.Exec("INSERT INTO members VALUES ($1, $2)", login, roomId)
	if err != nil {
		log.Printf("postgres/room - JoinRoom: %v\n", err)
		return err
	}
	return nil
}

func (r *RoomRepository) isRoomExists(id string) (is bool, err error) {
	err = r.db.QueryRow("SELECT id FROM rooms WHERE id=$1", id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		log.Printf("isRoomExists: %s\n", err)
		return false, err
	}
	return true, nil
}
