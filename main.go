package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaya-1302/PaymentAPI/controller"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/login", controller.AuthController).Methods("POST")
	r.HandleFunc("/customer", controller.GetLoggedInCustomerHandler).Methods("GET")
	r.HandleFunc("/logout", controller.LogoutHandler).Methods("POST")

	// Start the server
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
