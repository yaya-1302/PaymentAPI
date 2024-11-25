package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yaya-1302/PaymentAPI/model"
	"github.com/yaya-1302/PaymentAPI/service"
	"github.com/yaya-1302/PaymentAPI/utils"
)

func InitiatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var paymentReq model.PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&paymentReq)
	if err != nil {
		utils.FailureResponse(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	err = service.InitiatePayment(paymentReq.MerchantID, paymentReq.Amount)
	if err != nil {
		utils.FailureResponse(w, http.StatusBadRequest, fmt.Sprintf("Payment failed: %v", err))
		return
	}

	utils.SuccessResponse(w, "Payment successful", nil)
}
