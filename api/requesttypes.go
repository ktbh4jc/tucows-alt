package main

type OrderRequest struct {
	CustomerId int `json:"customerId"`
	ProductId  int `json:"productId"`
	Count      int `json:"count"`
	Intent     int `json:"intent"`
}

func NewOrderRequest(customerId int, productId int, count int) *OrderRequest {
	return &OrderRequest{
		CustomerId: customerId,
		ProductId:  productId,
		Count:      count,
	}
}
