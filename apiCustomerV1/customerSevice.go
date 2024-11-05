package apiCustomerV1

import (
	"encoding/json"
	"my-first-app/database"
	"my-first-app/entities"
	"my-first-app/models"
	"time"
)

// get all customer
func CustomerServiceGetAll() models.ResponseModel {

	db, err := database.ConnectDB()

	// check connection
	if err != nil {

		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.HttpCode = 400
		responseModel.Status = "fail"
		responseModel.Message = err.Error()
		responseModel.Timestamp = now.Unix()

		return responseModel
	}

	defer db.Close()

	// get data from table customer
	customers, err := entities.CustomerEntityGetAll(db)
	if err != nil {
		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.HttpCode = 400
		responseModel.Status = "fail"
		responseModel.Message = err.Error()
		responseModel.Timestamp = now.Unix()

		return responseModel
	}

	// after get data, send data to model
	var customerModels []CustomerModel
	for _, customer := range customers {
		var customerModel CustomerModel

		customerModel.Id = customer.Id
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

	rm, err := json.Marshal(customerModels)
	if err != nil {
		panic(err)
	}

	var responseModel models.ResponseModel
	now := time.Now()

	responseModel.HttpCode = 200
	responseModel.Status = "ok"
	responseModel.Message = "data found"
	responseModel.Timestamp = now.Unix()
	responseModel.Data = string(rm)

	return responseModel

}

func CustomerServiceAddNew(customerModel *CustomerModel) models.ResponseModel {

	db, err := database.ConnectDB()
	if err != nil {
		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.HttpCode = 400
		responseModel.Status = "fail"
		responseModel.Message = err.Error()
		responseModel.Timestamp = now.Unix()

		return responseModel
	}

	defer db.Close()

	var customerEntity entities.CustomerEntity
	customerEntity.CustomerName = customerModel.CustomerName
	customerEntity.CustomerGender = customerModel.CustomerGender
	customerEntity.CustomerIdentityNumber = customerModel.CustomerIdentityNumber
	customerEntity.CustomerBirthPlace = customerModel.CustomerBirthPlace
	customerEntity.CustomerBirthDate = customerModel.CustomerBirthDate
	var customerIdInserted int64 = entities.CustomerEntityInsert(db, &customerEntity)

	if customerIdInserted == 0 {
		var responseModel models.ResponseModel
		now := time.Now()

		responseModel.HttpCode = 400
		responseModel.Status = "fail"
		responseModel.Message = "no data found"
		responseModel.Timestamp = now.Unix()

		return responseModel
	}

	var responseModel models.ResponseModel
	now := time.Now()

	var data map[string]int64
	data = map[string]int64{}
	data["customerId"] = customerIdInserted
	jData, err := json.Marshal(data["customerId"])
	if err != nil {
		panic(err)
	}

	responseModel.HttpCode = 400
	responseModel.Status = "fail"
	responseModel.Message = "success"
	responseModel.Timestamp = now.Unix()
	responseModel.Data = string(jData)

	return responseModel
}
