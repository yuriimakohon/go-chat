package handler

import (
	"github.com/gin-gonic/gin"
)

var (
	msgBadCredsFormat     = "invalid credentials format"
	msgBadCredsJSONFormat = "invalid credentials JSON-format"
	msgBadCreds           = "wrong login or password"
	msgUserAlreadyExists  = "user already exists"
	msgUserDoesntExists   = "user doesn't exists"

	msgBadRoomFormat     = "invalid room info format"
	msgRoomAlreadyExists = "room already exists"
	msgRoomDoesntExists  = "room doesn't exists"
	msgUserAlreadyInRoom = "user already in the room"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"message": message})
}
