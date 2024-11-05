package models

type ResponseModel struct {
	HttpCode  int    `json:"-"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timeStamp"`
	Data      string `json:"data"`
}
