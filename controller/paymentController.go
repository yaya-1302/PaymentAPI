package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yaya-1302/PaymentAPI/model"
	"github.com/yaya-1302/PaymentAPI/service"
	"github.com/yaya-1302/PaymentAPI/utils"
)

func PaymentController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.FailureResponse(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	var request model.PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.FailureResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if request.MerchantID == "" || request.Amount <= 0 {
		utils.FailureResponse(w, http.StatusBadRequest, "Merchant ID and Amount must be provided and valid")
		return
	}

	err = service.InitiatePayment(request.MerchantID, request.Amount)
	if err != nil {
		utils.FailureResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, "Payment successful", nil)
}
