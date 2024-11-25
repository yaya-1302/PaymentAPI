package service

import (
	"errors"

	"github.com/yaya-1302/PaymentAPI/model"
	"github.com/yaya-1302/PaymentAPI/repository"
	"github.com/yaya-1302/PaymentAPI/utils"
)

var loggedInCustomer *model.Customer

func Login(username, password string) (string, error) {
	customers, err := repository.ReadCustomers()
	if err != nil {
		return "", err
	}

	for _, customer := range customers {
		if customer.Username == username && customer.Password == password {
			tokenString, err := utils.GenerateJWT(customer.Username)
			if err != nil {
				return "", err
			}

			loggedInCustomer = &customer

			return tokenString, nil
		}
	}

	return "", errors.New("invalid credentials")
}

func GetLoggedInCustomer() *model.Customer {
	return loggedInCustomer
}

func Logout() error {
	loggedInCustomer = nil
	return nil
}
