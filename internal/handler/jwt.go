package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/yuriimakohon/go-chat/config"
	"time"
)

var jwtSecret = []byte("Gw1_qzAm5")

func newToken() (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(config.TokenMaxAge).Unix(),
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)

	return tokenString, err
}
