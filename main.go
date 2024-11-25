package main

import (
	"log"
	"net/http"

	"github.com/yaya-1302/PaymentAPI/routes"
)

func main() {
	routes.RegisteredRoutes()

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
