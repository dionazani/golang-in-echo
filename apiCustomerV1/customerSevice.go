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
		now := time.Now()
		var responseModel models.ResponseModel = models.SetResponseModel(400, "fail", err.Error(), now.Unix(), "")
		return responseModel
	}

	defer db.Close()

	// get data from table customer
	customers, err := entities.CustomerEntityGetAll(db)
	if err != nil {
		now := time.Now()
		var responseModel models.ResponseModel = models.SetResponseModel(400, "fail", err.Error(), now.Unix(), "")
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

	now := time.Now()
	var responseModel models.ResponseModel = models.SetResponseModel(200, "ok", "data found", now.Unix(), string(rm))

	return responseModel
}

// add new customer
func CustomerServiceAddNew(customerModel *CustomerModel) models.ResponseModel {

	// check database connection.
	db, err := database.ConnectDB()
	if err != nil {
		now := time.Now()
		var responseModel models.ResponseModel = models.SetResponseModel(400, "fail", err.Error(), now.Unix(), "")
		return responseModel
	}

	defer db.Close()

	// insert new customer
	var customerEntity entities.CustomerEntity
	customerEntity.CustomerName = customerModel.CustomerName
	customerEntity.CustomerGender = customerModel.CustomerGender
	customerEntity.CustomerIdentityNumber = customerModel.CustomerIdentityNumber
	customerEntity.CustomerBirthPlace = customerModel.CustomerBirthPlace
	customerEntity.CustomerBirthDate = customerModel.CustomerBirthDate
	var customerIdInserted int64 = entities.CustomerEntityInsert(db, &customerEntity)

	// check customerId.
	if customerIdInserted == 0 {

		now := time.Now()
		var responseModel models.ResponseModel = models.SetResponseModel(400, "fail", "cannot insert new customer", now.Unix(), "")

		return responseModel
	}

	var data = map[string]int64{}
	data["customerId"] = customerIdInserted
	jData, err := json.Marshal(data["customerId"])
	if err != nil {
		data["customerId"] = 0
	}

	now := time.Now()
	var responseModel models.ResponseModel = models.SetResponseModel(201, "created", "success", now.Unix(), string(jData))

	return responseModel
}

// add new customer
func CustomerServiceUpdate(customerModel *CustomerModel) models.ResponseModel {

	// check database connection.
	db, err := database.ConnectDB()
	if err != nil {
		now := time.Now()
		var responseModel models.ResponseModel = models.SetResponseModel(400, "fail", err.Error(), now.Unix(), "")
		return responseModel
	}

	defer db.Close()

	// update customer.
	var customerEntity entities.CustomerEntity
	customerEntity.Id = customerModel.Id
	customerEntity.CustomerName = customerModel.CustomerName
	customerEntity.CustomerGender = customerModel.CustomerGender
	customerEntity.CustomerIdentityNumber = customerModel.CustomerIdentityNumber
	customerEntity.CustomerBirthPlace = customerModel.CustomerBirthPlace
	customerEntity.CustomerBirthDate = customerModel.CustomerBirthDate

	now := time.Now()
	var responseModel models.ResponseModel = models.SetResponseModel(200, "success", "success", now.Unix(), "")

	return responseModel
}
