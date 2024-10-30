package models

type ResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
