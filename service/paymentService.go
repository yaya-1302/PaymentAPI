package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yaya-1302/PaymentAPI/model"
	"github.com/yaya-1302/PaymentAPI/repository"
)

func InitiatePayment(merchantID string, amount float64) error {
	customer, err := validateCustomer()
	if err != nil {
		return err
	}

	merchant, err := validateMerchant(merchantID)
	if err != nil {
		return err
	}

	if customer.Balance < amount {
		return errors.New("insufficient funds")
	}

	customer.Balance -= amount

	merchant.Balance += amount

	if err := repository.UpdateCustomerBalance(customer); err != nil {
		return fmt.Errorf("failed to update customer balance: %v", err)
	}

	if err := repository.UpdateMerchantBalance(merchant); err != nil {
		return fmt.Errorf("failed to update merchant balance: %v", err)
	}

	err = addPaymentHistory(customer, amount)
	if err != nil {
		return fmt.Errorf("failed to add payment history: %v", err)
	}

	return nil
}

func validateCustomer() (*model.Customer, error) {
	customer := GetLoggedInCustomer()

	if customer == nil {
		return nil, errors.New("customer not logged in")
	}

	return customer, nil
}

func validateMerchant(merchantID string) (*model.Merchant, error) {
	merchants, err := repository.ReadMerchants()
	if err != nil {
		return nil, err
	}

	for _, merchant := range merchants {
		if merchant.ID == merchantID {
			return &merchant, nil
		}
	}

	return nil, errors.New("merchant not found")
}

func addPaymentHistory(customer *model.Customer, amount float64) error {
	historyID := uuid.New().String()

	history := model.History{
		ID:         historyID,
		CustomerID: customer.ID,
		Action:     "payment",
		Amount:     amount,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	return repository.AddHistory(history)
}
