package model

import "golang.org/x/crypto/bcrypt"

type Customer struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

func (customer *Customer) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	customer.Password = string(hashedPassword)
	return nil
}

func (customer *Customer) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	return err == nil
}
