package models

type ResponseModel struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timeStamp"`
	Data      string `json:"data"`
}
