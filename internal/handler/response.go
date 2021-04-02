package handler

import (
	"github.com/gin-gonic/gin"
)

var (
	msgBadCredsFormat        = "invalid credentials format"
	msgBadCreds              = "wrong login or password"
	msgUserAlreadyExists     = "user already exists"
	msgInvalidTokenSignature = "invalid token's signature"
	msgInvalidToken          = "invalid token's signature"
	msgNoTokenCookie         = `'token' cookie not present`
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"message": message})
}
