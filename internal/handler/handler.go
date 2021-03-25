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

func (h Handler) renderHTML(pathname string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, pathname, nil)
	}
}

func (h *Handler) SetupRoutes() http.Handler {
	r := gin.Default()
	r.LoadHTMLGlob(config.HTMLPath)

	auth := r.Group("/auth")
	{
		auth.GET("signup", h.renderHTML(config.SignupPage))
		auth.POST("signup", h.signupHandler, h.renderToken)

		auth.GET("login", h.renderHTML(config.LoginPage))
		auth.POST("login", h.loginHandler, h.renderToken)
	}

	r.GET("/", h.authRequired(), h.renderHTML(config.IndexPage))
	r.GET("/ws", h.wsHandler)

	return r
}
