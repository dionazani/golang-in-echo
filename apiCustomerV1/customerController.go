package apiCustomerV1

import (
	"my-first-app/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// function get all customer
func CustomerControllerGetAll(c echo.Context) error {

	var responseModel models.ResponseModel = CustomerServiceGetAll()
	return c.JSON(responseModel.HttpCode, responseModel)
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

	var responseModel models.ResponseModel = CustomerServiceAddNew(customerModel)
	return c.JSON(responseModel.HttpCode, responseModel)
}
