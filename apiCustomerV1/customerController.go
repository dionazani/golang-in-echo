package apiCustomerV1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"my-first-app/models"
)

// function get all customer
func CustomerControllerGetAll(c echo.Context) error {

	customerModels, responseCode, message := CustomerServiceGetAll()

	if customerModels == nil && responseCode != 200 {
		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.Status = "fail"
		responseModel.Message = message
		responseModel.Timestamp = now.Unix()

		return c.JSON(responseCode, responseModel)
	}

	out, err := json.Marshal(customerModels)
	if err != nil {
		panic(err)
	}

	var responseModel models.ResponseModel
	now := time.Now()

	responseModel.Status = "ok"
	responseModel.Message = "success"
	responseModel.Timestamp = now.Unix()
	responseModel.Data = string(out[:])

	return c.JSON(responseCode, responseModel)
}

func CustomerControllerAddNew(c echo.Context) error {

	customerModel := new(CustomerModel)

	if err := c.Bind(customerModel); err != nil {
		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.Status = "bad request"
		responseModel.Message = err.Error()
		responseModel.Timestamp = now.Unix()

		return c.JSON(http.StatusBadRequest, responseModel)
	}

	customerModels, responseCode, message := CustomerServiceAddNew(customerModel)

	if customerModels == nil && responseCode != 200 {
		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.Status = "fail"
		responseModel.Message = message
		responseModel.Timestamp = now.Unix()

		return c.JSON(responseCode, responseModel)
	}

	var responseModel models.ResponseModel
	now := time.Now()

	responseModel.Status = "created"
	responseModel.Message = "success"
	responseModel.Timestamp = now.Unix()

	return c.JSON(http.StatusCreated, responseModel)
}
