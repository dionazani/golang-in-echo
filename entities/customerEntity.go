package entities

import "database/sql"

type CustomerEntity struct {
	Id                     int32
	CustomerName           string
	CustomerGender         string
	CustomerIdentityNumber string
	CustomerBirthPlace     string
	CustomerBirthDate      string
}

func CustomerEntityGetAll(db *sql.DB) ([]CustomerEntity, error) {

	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []CustomerEntity

	for rows.Next() {
		var customer CustomerEntity
		if err := rows.Scan(&customer.Id, &customer.CustomerName, &customer.CustomerGender, &customer.CustomerIdentityNumber, &customer.CustomerBirthPlace, &customer.CustomerBirthDate); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil

}

func CustomerEntityInsert(db *sql.DB, customerEntity *CustomerEntity) int64 {

	var customerIdInserted int64 = 0

	res, err := db.Exec("INSERT INTO customer (customer_name,customer_gender,customer_identity_number,customer_birth_place,customer_birth_date) VALUES (?, ?, ?, ?, ?)",
		customerEntity.CustomerName,
		customerEntity.CustomerGender,
		customerEntity.CustomerIdentityNumber,
		customerEntity.CustomerBirthPlace,
		customerEntity.CustomerBirthDate)

	if err != nil {
		customerIdInserted = 0
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			customerIdInserted = 0
		} else {
			customerIdInserted = id
		}
	}

	return customerIdInserted
}

func CustomerEntityUpdate(db *sql.DB, customerEntity *CustomerEntity) int64 {

	var customerUpdated int64 = 0

	_, err := db.Exec("UPDATE customer set customer_name = ?,customer_gender=?,customer_identity_number=?,customer_birth_place=?,customer_birth_date = ? WHERE id=?",
		customerEntity.CustomerName,
		customerEntity.CustomerGender,
		customerEntity.CustomerIdentityNumber,
		customerEntity.CustomerBirthPlace,
		customerEntity.CustomerBirthDate,
		customerEntity.Id)

	if err != nil {
		customerUpdated = 0
	} else {
		customerUpdated = 1
	}

	return customerUpdated
}
