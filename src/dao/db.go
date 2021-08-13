package dao

import (
	"database/sql"
	"dto"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connectToDataBase() (*sql.DB, error) {
	return sql.Open("mysql", "saber66:AdminSaber66@tcp(localhost:3306)/test")
}

func FindAllCustomers() ([]dto.Customer, error) {

	db, err := connectToDataBase()
	if err != nil {
		return nil, errors.New("cannot connect to database " + err.Error())
	}
	var customers []dto.Customer

	customer := dto.Customer{}
	rows, err := db.Query("select * from customers")
	if err != nil {
		return nil, errors.New("cannot fetch data from database" + err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Email)
		if err != nil {
			return nil, errors.New("cannot fetch data from database" + err.Error())
		}
		customers = append(customers, customer)
	}
	return customers, nil

}

func FindCustomerById(id int) (*dto.Customer, error) {

	db, err := connectToDataBase()
	var customer dto.Customer
	if err != nil {
		return nil, errors.New("cannot connect to database " + err.Error())
	}

	statement, err := db.Prepare("select * from customers where id=?")
	if err != nil {
		return nil, errors.New("cannot fetch data from database" + err.Error())
	}
	rows, err := statement.Query(id)
	if err != nil {
		return nil, errors.New("cannot fetch data from database" + err.Error())
	}
	if !rows.Next() {
		return nil, fmt.Errorf("customer does not exist with id %d", id)
	} else {
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Email)
		if err != nil {
			return nil, errors.New("cannot fetch data from database" + err.Error())
		}
	}
	return &customer, nil

}
func AddCustomer(customer dto.Customer) (bool, error) {
	db, err := connectToDataBase()
	if err != nil {
		return false, fmt.Errorf("cannot connect to database %v", err.Error())
	}
	statement, err := db.Prepare("insert into customers (first_name, last_name, email) values (?,?,?)")

	if err != nil {
		return false, fmt.Errorf("cannot insert to  database %v", err.Error())
	}
	result, err := statement.Exec(customer.FirstName, customer.LastName, customer.Email)
	if err != nil {
		return false, fmt.Errorf("cannot insert to  database %v", err.Error())
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("cannot insert to  database %v", err.Error())
	}
	if rowAffected > 0 {
		return true, nil
	} else {
		return false, errors.New("cannot insert to  database ")
	}
}

func UpdateCustomer(customer dto.Customer, id int) (bool, error) {
	db, err := connectToDataBase()
	if err != nil {
		return false, fmt.Errorf("cannot connect to database %v", err.Error())
	}
	statement, err := db.Prepare("update customers set first_name=? , last_name=? , email=? where id=?")

	if err != nil {
		return false, fmt.Errorf("cannot update table customer %v", err.Error())
	}
	result, err := statement.Exec(customer.FirstName, customer.LastName, customer.Email, id)

	if err != nil {
		return false, fmt.Errorf("cannot update table customer %v", err.Error())
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("cannot update table customer %v", err.Error())
	}
	if rowAffected > 0 {
		return true, nil
	} else {
		return false, errors.New("customer does not exist ")
	}
}

func DeleteCustomer(id int) (bool, error) {
	db, err := connectToDataBase()
	if err != nil {
		return false, fmt.Errorf("cannot connect to database %v", err.Error())
	}
	statement, err := db.Prepare("delete from customers  where id=?")

	if err != nil {
		return false, fmt.Errorf("cannot delete row table  customer %v", err.Error())
	}

	result, err := statement.Exec(id)

	if err != nil {
		return false, fmt.Errorf("cannot delete row table  customer %v", err.Error())
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("cannot delete row table  customer %v", err.Error())
	}
	if rowAffected > 0 {
		return true, nil
	} else {
		return false, errors.New("customer does not exist ")
	}

}
