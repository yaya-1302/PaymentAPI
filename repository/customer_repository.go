package repository

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/yaya-1302/PaymentAPI/model"
)

func GetCustomers() ([]model.Customer, error) {
	file, err := os.Open("data/customers.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var customers []model.Customer
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&customers)
	return customers, err
}

func GetCustomerByUsername(username string) (*model.Customer, error) {
	customers, err := GetCustomers()
	if err != nil {
		return nil, err
	}
	for _, customer := range customers {
		if customer.Username == username {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
