package handler

import (
	"github.com/gin-gonic/gin"
	w "github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var wu = w.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) wsHandler(ctx *gin.Context) {
	conn, err := wu.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	defer conn.Close()

	log.Println("User connected")

	stop := make(chan interface{})
	conn.SetCloseHandler(func(code int, text string) error {
		log.Println("User disconnected: ", text)
		stop <- struct{}{}
		return nil
	})

	for {
		select {
		case <-stop:
			break
		default:
			t, data, err := conn.ReadMessage()
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}

			err = conn.WriteMessage(t, []byte(strings.ToUpper(string(data))))
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}
	}
}
