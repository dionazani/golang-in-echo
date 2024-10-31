package apiCustomer

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"my-first-app/database"
	"my-first-app/entities"
	"my-first-app/models"
)

// function get all customer
func GetAll(c echo.Context) error {
	db, err := database.ConnectDB()
	if err != nil {

		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.Status = "fail"
		responseModel.Message = "connection invalid"
		responseModel.Timestamp = now.Unix()

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "connection failed"})

	}
	defer db.Close()

	customers, err := entities.GetAllCustomers(db)
	if err != nil {

		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.Status = "fail"
		responseModel.Message = "customer loading data failed"
		responseModel.Timestamp = now.Unix()

		return c.JSON(http.StatusInternalServerError, responseModel)
	}

	var customerModels []CustomerModel

	if len(customers) > 0 {
		for _, customer := range customers {
			var customerModel CustomerModel

			customerModel.CustomerName = customer.CustomerName
			if customer.CustomerGender == "P" {
				customerModel.CustomerGender = "Pria"
			} else {
				customerModel.CustomerGender = "Wanita"
			}
			customerModel.CustomerIdentityNumber = customer.CustomerIdentityNumber
			customerModel.CustomerBirthPlace = customer.CustomerBirthPlace
			customerModel.CustomerBirthDate = customer.CustomerBirthDate

			customerModels = append(customerModels, customerModel)
		}
	} else {
		customerModels = nil
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

	return c.JSON(http.StatusOK, responseModel)
}

func AddNew(c echo.Context) error {
	db, err := database.ConnectDB()
	if err != nil {

		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.Status = "error"
		responseModel.Message = "connection fail"
		responseModel.Timestamp = now.Unix()

		return c.JSON(http.StatusInternalServerError, responseModel)
	}
	defer db.Close()

	customerModel := new(CustomerModel)
	if err := c.Bind(customerModel); err != nil {

		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.Status = "error"
		responseModel.Message = err.Error()
		responseModel.Timestamp = now.Unix()

		return c.JSON(http.StatusBadRequest, responseModel)
	}

	var customerEntity entities.CustomerEntity
	customerEntity.CustomerName = customerModel.CustomerName
	customerEntity.CustomerGender = customerModel.CustomerGender
	customerEntity.CustomerIdentityNumber = customerModel.CustomerIdentityNumber
	customerEntity.CustomerBirthPlace = customerModel.CustomerBirthPlace
	customerEntity.CustomerBirthDate = customerModel.CustomerBirthDate

	err = entities.InsertCustomer(db, &customerEntity)
	if err != nil {

		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.Status = "error"
		responseModel.Message = err.Error()
		responseModel.Timestamp = now.Unix()
		return c.JSON(http.StatusInternalServerError, responseModel)
	}

	var responseModel models.ResponseModel
	now := time.Now()

	responseModel.Status = "created"
	responseModel.Message = "success"
	responseModel.Timestamp = now.Unix()

	return c.JSON(http.StatusCreated, responseModel)
}
