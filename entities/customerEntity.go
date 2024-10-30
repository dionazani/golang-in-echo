package entities

import "database/sql"

type Customer struct {
	Id                     int32
	CustomerName           string
	CustomerGender         string
	CustomerIdentityNumber string
}

func GetAllCustomers(db *sql.DB) ([]Customer, error) {

	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.Id, &customer.CustomerName, &customer.CustomerGender, &customer.CustomerIdentityNumber); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil

}
