package model

type History struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	Action     string  `json:"action"`
	Amount     float64 `json:"amount"`
	Timestamp  string  `json:"timestamp"`
}
