package service

import (
	"github.com/spf13/viper"
	"github.com/yuriimakohon/go-chat/internal/models"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(creds models.Credentials) error {
	if !isValidCredsFormat(creds) {
		return ErrBadCredsFormat
	}

	hashedPassword, err := generatePasswordHash([]byte(creds.Password))
	if err != nil {
		return err
	}

	creds.Password = string(hashedPassword)
	if err = s.repo.CreateUser(creds); err != nil {
		if err == repository.ErrUserAlreadyExists {
			return ErrUserAlreadyExists
		}
		return err
	}

	return nil
}

func (s *AuthService) LogIn(creds models.Credentials) error {
	if !isValidCredsFormat(creds) {
		return ErrBadCredsFormat
	}

	storedCreds, err := s.repo.GetUser(creds.Login)
	if err != nil {
		if err == repository.ErrUserDoesntExists {
			return ErrUserDoesntExists
		}
		log.Printf("service-auth: %s\n", err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ErrWrongPassword
		}
		log.Printf("service-auth: %s\n", err)
		return err
	}

	return nil
}

func generatePasswordHash(password []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		viper.GetInt("password_hash_cost"))
	if err != nil {
		log.Printf("service-auth - generatePassordHash: %s\n", err)
		return nil, err
	}
	return hashedPassword, nil
}

var (
	loginRegexp = regexp.MustCompile("^[A-Za-z0-9_-]{3,30}$")
	passRegexp  = regexp.MustCompile("^[A-Za-z0-9_-]{5,255}$")
)

func isValidCredsFormat(creds models.Credentials) bool {
	if !loginRegexp.MatchString(creds.Login) ||
		!passRegexp.MatchString(creds.Password) {
		return false
	}
	return true
}
