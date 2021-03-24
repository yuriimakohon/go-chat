package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yuriimakohon/go-chat/config"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"net/http"
)

type Handler struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) SetupRoutes() http.Handler {
	r := gin.Default()
	r.LoadHTMLGlob(config.HTMLPath)

	auth := r.Group("/auth")
	{
		auth.GET("signup", h.signupGetHandler)
		auth.POST("signup", h.signupHandler)

		auth.GET("login", h.loginGetHandler)
		auth.POST("login", h.loginHandler)
	}

	r.GET("/", h.authRequired(), func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, config.IndexPage, nil)
	})
	r.GET("/ws", h.wsHandler)

	return r
}
