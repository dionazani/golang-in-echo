package models

type ResponseModel struct {
	HttpCode  int    `json:"-"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timeStamp"`
	Data      string `json:"data"`
}

func SetResponseModel(httpCode int,
	status string,
	message string,
	timestamp int64,
	data string) ResponseModel {

	var responseModel ResponseModel
	responseModel.HttpCode = httpCode
	responseModel.Status = status
	responseModel.Message = message
	responseModel.Timestamp = timestamp
	responseModel.Data = data

	return responseModel
}
