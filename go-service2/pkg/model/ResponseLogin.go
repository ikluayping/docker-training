package model

type ResponseLogin struct {
	ResponseCode string `json:"status,omitempty"`
	ResponseDesc string `json:"description,omitempty"`
}
