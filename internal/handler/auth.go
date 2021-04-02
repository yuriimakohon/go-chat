package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yuriimakohon/go-chat/internal/models"
	"github.com/yuriimakohon/go-chat/internal/service"
	"log"
	"net/http"
	"os"
)

func (h *Handler) logIn(c *gin.Context) {
	creds := models.Credentials{}
	if err := c.ShouldBindJSON(&creds); err != nil {
		newErrorResponse(c, http.StatusBadRequest, msgBadCredsFormat)
		return
	}

	if err := h.service.LogIn(creds); err != nil {
		if err == service.ErrWrongPassword || err == service.ErrUserDosentExists {
			log.Printf("Wrong credentials: %v", err)
			newErrorResponse(c, http.StatusBadRequest, msgBadCreds)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) signUp(c *gin.Context) {
	creds := models.Credentials{}
	if err := c.ShouldBindJSON(&creds); err != nil {
		newErrorResponse(c, http.StatusBadRequest, msgBadCredsFormat)
		return
	}

	if err := h.service.SignUp(creds); err != nil {
		if err == service.ErrUserAlreadyExists {
			newErrorResponse(c, http.StatusConflict, msgUserAlreadyExists)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) setTokenCookieMiddleware(c *gin.Context) {
	token := jwt.New(jwt.SigningMethodHS256)
	signedTokenStr, err := token.SignedString([]byte(os.Getenv("TOKEN_SIGN")))
	if err != nil {
		log.Printf("setTokenCookieMiddleware: %s\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.SetCookie(
		"token", signedTokenStr, 7,
		"", "", false, false,
	)
}

func (h *Handler) authRequiredMiddleware(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusFound, "/auth/login")
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SIGN")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
			return
		}
		log.Printf("authRequiredMiddleware: %s\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !token.Valid {
		c.Redirect(http.StatusFound, "/auth/login")
		c.Abort()
		return
	}
}
