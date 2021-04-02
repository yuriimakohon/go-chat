package models

type Credentials struct {
	Login    string `json:"login" db:"login" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
