package model

type SuccessMessage struct {
	Errcode string      `json:"errcode,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrMessage struct {
	Errcode string `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
