package credentials

type Credentials struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}
