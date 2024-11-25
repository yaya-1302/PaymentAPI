package routes

import (
	"net/http"

	"github.com/yaya-1302/PaymentAPI/controller"
)

func RegisteredRoutes() {
	http.HandleFunc("/login", controller.LoginController)
	http.HandleFunc("/logout", controller.LogoutController)
	http.HandleFunc("/payment", controller.PaymentController)
}
