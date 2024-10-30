package apiCustomer

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"my-first-app/database"
	"my-first-app/entities"
	"my-first-app/models"
)

func GetAllCustomers(c echo.Context) error {
	// Mengambil koneksi database
	db, err := database.ConnectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "connection failed"})
	}
	defer db.Close()

	customers, err := entities.GetAllCustomers(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "customer loading data failed"})
	}

	out, err := json.Marshal(customers)
	if err != nil {
		panic(err)
	}

	var responseModel models.ResponseModel
	responseModel.Status = "200"
	responseModel.Message = "success"
	responseModel.Data = string(out[:])

	return c.JSON(http.StatusOK, responseModel)
}
