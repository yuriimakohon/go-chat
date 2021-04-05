package models

type Room struct {
	Id string `json:"id" binding:"required"`
}
