package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yuriimakohon/go-chat/config"
	"github.com/yuriimakohon/go-chat/internal/models/credentials"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"log"
	"net/http"
)

func (h *Handler) signupGetHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, config.SignupPage, nil)
}

func (h *Handler) loginGetHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, config.LoginPage, nil)
}

func (h *Handler) signupHandler(ctx *gin.Context) {
	creds := credentials.Credentials{}
	if json.NewDecoder(ctx.Request.Body).Decode(&creds) != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.repo.NewUser(creds); err != nil {
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

	creds, err := h.repo.GetUserByLogin(creds.Login)
	if err != nil {
		if err == repository.ErrUserNotFound {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, creds)
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
