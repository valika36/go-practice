package Models

type TransactionResponse struct {
	OrderID    string `json:"orderId"`
	CustomerId string `json:"customerId"`
	ProductId  string `json:"productId"`
	Quantity   uint   `json:"quantity"`
	TotalPrice uint   `json:"totalPrice"`
	Status     string `json:"status"`
}
