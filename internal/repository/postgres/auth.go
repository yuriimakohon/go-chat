package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/yuriimakohon/go-chat/internal/models"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"log"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUser(login string) (creds models.Credentials, err error) {
	err = r.db.QueryRow("SELECT * FROM users WHERE login=$1", login).Scan(&creds.Login, &creds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return creds, repository.ErrUserDoesntExists
		}
		log.Printf("postgres/auth - GetUser: %s\n", err)
	}
	return creds, err
}

func (r *AuthRepository) CreateUser(creds models.Credentials) (err error) {
	if err = r.isUserExists(creds.Login); err != nil {
		return err
	}

	_, err = r.db.Exec("INSERT INTO users VALUES ($1, $2)", creds.Login, creds.Password)
	if err != nil {
		log.Printf("postgres/auth - CreateUser: %s\n", err)
		return err
	}
	return nil
}

func (r *AuthRepository) isUserExists(login string) (err error) {
	err = r.db.QueryRow("SELECT login FROM users WHERE login=$1", login).Scan(&login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Printf("isUserExists: %s\n", err)
		return err
	}
	return repository.ErrUserAlreadyExists
}
