package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	apiCustomer "my-first-app/apiCustomerV1"
	"my-first-app/database"
)

func main() {

	// Koneksi ke database MySQL
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Gagal terkoneksi ke database:", err)
	}
	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/customers", apiCustomer.CustomerControllerGetAll)
	e.POST("/customers", apiCustomer.CustomerControllerAddNew)

	e.Logger.Fatal(e.Start(":9031"))
}
