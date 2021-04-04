package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuriimakohon/go-chat/internal/models"
	"github.com/yuriimakohon/go-chat/internal/service"
	"log"
	"net/http"
)

func (h *Handler) joinRoom(c *gin.Context) {
	login, ok := getUserLogin(c)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	roomId := c.Param("id")

	if err := h.service.JoinRoom(login, roomId); err != nil {
		switch err {
		case service.ErrRoomDosentExists:
			newErrorResponse(c, http.StatusNotFound, msgRoomDosentExists)
		case service.ErrUserAlreadyInRoom:
			newErrorResponse(c, http.StatusConflict, msgUserAlreadyInRoom)
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
}

func (h *Handler) createRoom(c *gin.Context) {
	room := models.Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		newErrorResponse(c, http.StatusBadRequest, msgBadRoomFormat)
		return
	}

	if err := h.service.CreateRoom(room); err != nil {
		if err == service.ErrRoomAlreadyExists {
			newErrorResponse(c, http.StatusConflict, msgRoomAlreadyExists)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//h.joinRoom(c)
}

func (h *Handler) getRooms(c *gin.Context) {
	login, ok := getUserLogin(c)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	panic(fmt.Sprintf("Implement me: getRooms for user: %s\n", login))
}

func (h *Handler) connectRoom(c *gin.Context) {
	login, ok := getUserLogin(c)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	roomId := c.Param("id")

	panic(fmt.Sprintf("Implement me: connectRoom to id: %s | user: %s\n", roomId, login))
}

func getUserLogin(c *gin.Context) (string, bool) {
	loginI, ok := c.Get("userLogin")
	if !ok {
		log.Println("getUserLogin: no 'userLogin' var in ctx")
		return "", false
	}
	loginStr, ok := loginI.(string)
	if !ok {
		log.Println("getUserLogin: 'userLogin' assert - not a string")
		return "", false
	}
	return loginStr, true
}
