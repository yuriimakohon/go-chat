package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yuriimakohon/go-chat/config"
	"github.com/yuriimakohon/go-chat/internal/models/credentials"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func (h *Handler) signupHandler(ctx *gin.Context) {
	creds := credentials.Credentials{}
	if json.NewDecoder(ctx.Request.Body).Decode(&creds) != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), config.BcryptCost)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	creds.Password = string(hashedPassword)

	if err = h.repo.NewUser(creds); err != nil {
		if err == repository.ErrUserAlreadyExists {
			ctx.AbortWithStatusJSON(http.StatusConflict,
				gin.H{"message": "User already exists, choose another login"})
			return
		}
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) loginHandler(ctx *gin.Context) {
	creds := credentials.Credentials{}
	if json.NewDecoder(ctx.Request.Body).Decode(&creds) != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	storedCreds, err := h.repo.GetUserByLogin(creds.Login)
	if err != nil {
		if err == repository.ErrUserNotFound {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"message": "Wrong login or password"})
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)) != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Wrong login or password"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successful login"})
}

func (h *Handler) authRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.Query("t")
		if t == "" {
			ctx.Redirect(http.StatusFound, "/auth/login")
			ctx.Abort()
		}
	}
}
