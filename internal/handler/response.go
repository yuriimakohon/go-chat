package handler

import (
	"github.com/gin-gonic/gin"
)

var (
	msgBadCredsFormat    = "invalid credentials format"
	msgBadCreds          = "wrong login or password"
	msgUserAlreadyExists = "user already exists"
	msgUserDosentExists  = "user dosen't exists"

	msgBadRoomFormat     = "invalid room info format"
	msgRoomAlreadyExists = "room already exists"
	msgRoomDosentExists  = "room dosen't exists"
	msgUserAlreadyInRoom = "user already in the room"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"message": message})
}
