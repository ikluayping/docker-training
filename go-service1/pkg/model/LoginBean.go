package model

type LoginBean struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
