package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dwh:1234@tcp(localhost:3306)/my_first_golang")
	if err != nil {
		return nil, err
	}

	// Memeriksa koneksi ke database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("successfully connected")
	return db, nil
}
