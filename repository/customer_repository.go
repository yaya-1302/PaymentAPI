package repository

import (
	"os"

	"github.com/yaya-1302/PaymentAPI/model"
	"github.com/yaya-1302/PaymentAPI/utils"
)

func ReadCustomers() ([]model.Customer, error) {
	var customers []model.Customer

	err := utils.ReadFromFile("data/customers.json", &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func WriteCustomers(customers []model.Customer) error {
	err := utils.WriteToFile("data/customers.json", customers)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCustomerBalance(updatedCustomer *model.Customer) error {
	customers, err := ReadCustomers()
	if err != nil {
		return err
	}

	for i, customer := range customers {
		if customer.ID == updatedCustomer.ID {
			customers[i].Balance = updatedCustomer.Balance
			break
		}
	}

	return WriteCustomers(customers)
}

func AddHistory(history model.History) error {
	var histories []model.History

	if _, err := os.Stat("data/history.json"); err == nil {
		err := utils.ReadFromFile("data/history.json", &histories)
		if err != nil {
			return err
		}
	}

	histories = append(histories, history)

	return utils.WriteToFile("data/history.json", histories)
}
