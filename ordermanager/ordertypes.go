package main

type Order struct {
	ID         int    `json:"id"`
	CustomerId int    `json:"customerId"`
	ProductId  int    `json:"productId"`
	Count      int    `json:"count"`
	Status     string `json:"status"`
}

func NewOrder(customerId int, productId int, count int) *Order {
	return &Order{
		CustomerId: customerId,
		ProductId:  productId,
		Count:      count,
		Status:     "INITIALIZED",
	}
}
