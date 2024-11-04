package apiCustomerV1

import (
	"my-first-app/database"
	"my-first-app/entities"
)

// get all customer
func CustomerServiceGetAll() (customerModelList []CustomerModel, responseCode int, message string) {

	db, err := database.ConnectDB()
	if err != nil {
		return nil, 500, "connection database failed"
	}

	defer db.Close()

	// get data from table customer
	customers, err := entities.CustomerEntityGetAll(db)
	if err != nil {
		return nil, 500, "load data failed"
	}

	// after get data, send data to model
	var customerModels []CustomerModel

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

	return customerModels, 200, "success"

}

func CustomerServiceAddNew(customerModel *CustomerModel) (customerModelList []CustomerModel, responseCode int, message string) {

	db, err := database.ConnectDB()
	if err != nil {
		return nil, 500, "connection database failed"
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
		return nil, 500, "error while add new customer"
	}

	return nil, 200, "invalid request"
}
