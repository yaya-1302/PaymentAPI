package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yaya-1302/PaymentAPI/service"
	"github.com/yaya-1302/PaymentAPI/utils"
)

func AuthController(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		utils.FailureResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	token, err := service.Login(requestData.Username, requestData.Password)
	if err != nil {
		utils.FailureResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(w, "Login successful", struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
}

func GetLoggedInCustomerHandler(w http.ResponseWriter, r *http.Request) {
	customer := service.GetLoggedInCustomer()
	if customer == nil {
		utils.FailureResponse(w, http.StatusUnauthorized, "No logged-in customer")
		return
	}

	utils.SuccessResponse(w, "Logged-in customer retrieved", customer)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	err := service.Logout()
	if err != nil {
		utils.FailureResponse(w, http.StatusInternalServerError, "Error logging out")
		return
	}

	utils.SuccessResponse(w, "Logout successful", nil)
}
