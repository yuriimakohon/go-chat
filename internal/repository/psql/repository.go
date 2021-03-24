package psql

import (
	_ "github.com/lib/pq"
	"github.com/yuriimakohon/go-chat/internal/models/credentials"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"log"
)
import "database/sql"

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	db, err := sql.Open("postgres", "dbname=postgres user=postgres password=postgres")
	if err != nil {
		log.Println("Cannot open psql db: ", err)
	}

	return &Repository{db: db}
}

func (r *Repository) NewUser(creds credentials.Credentials) error {
	_, err := r.GetUserByLogin(creds.Login)
	if err != nil {
		if err != repository.ErrUserNotFound {
			return err
		}
	} else {
		return repository.ErrUserAlreadyExists
	}

	_, err = r.db.Query("insert into users values ($1, $2)",
		creds.Login,
		creds.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByLogin(login string) (credentials.Credentials, error) {
	result := r.db.QueryRow("select * from users where login=$1", login)

	storedCreds := credentials.Credentials{}
	err := result.Scan(&storedCreds.Login, &storedCreds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return credentials.Credentials{}, repository.ErrUserNotFound
		}
		return credentials.Credentials{}, err
	}

	return storedCreds, nil
}
