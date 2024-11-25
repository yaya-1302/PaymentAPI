package repository

import (
	"github.com/yaya-1302/PaymentAPI/model"
	"github.com/yaya-1302/PaymentAPI/utils"
)

func ReadMerchants() ([]model.Merchant, error) {
	var merchants []model.Merchant

	err := utils.ReadFromFile("data/merchants.json", &merchants)
	if err != nil {
		return nil, err
	}

	return merchants, nil
}

func WriteMerchants(merchants []model.Merchant) error {
	err := utils.WriteToFile("data/merchants.json", merchants)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMerchantBalance(updatedMerchant *model.Merchant) error {
	merchants, err := ReadMerchants()
	if err != nil {
		return err
	}

	for i, merchant := range merchants {
		if merchant.ID == updatedMerchant.ID {
			merchants[i].Balance = updatedMerchant.Balance
			break
		}
	}

	return WriteMerchants(merchants)
}
