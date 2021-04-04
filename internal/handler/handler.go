package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yuriimakohon/go-chat/configs"
	"github.com/yuriimakohon/go-chat/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) renderHTML(pathname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, pathname, nil)
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob(configs.HTMLPath)

	auth := r.Group("/auth")
	{
		auth.GET("/login", h.renderHTML(configs.LoginPage))
		auth.GET("/signup", h.renderHTML(configs.SignupPage))

		auth.POST("/login", h.logIn, h.setTokenCookieMiddleware)
		auth.POST("/signup", h.signUp, h.setTokenCookieMiddleware)
	}

	room := r.Group("room")
	room.Use(h.authRequiredMiddleware, h.setTokenCookieMiddleware)
	{
		room.POST("/", h.createRoom)
		room.GET("/", h.getRooms)
		room.POST("/:id", h.joinRoom)
		room.GET("/:id", h.connectRoom)
	}

	msg := r.Group("/msg")
	{
		msg.POST("/:room_id", h.createMsg)
	}

	r.GET("/", h.authRequiredMiddleware, h.setTokenCookieMiddleware, h.renderHTML(configs.IndexPage))

	return r
}
