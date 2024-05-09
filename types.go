// A place to store all of the types used in this project.

package main

type OrderRequest struct {
	CustomerId int `json:"customerId"`
	ProductId  int `json:"productId"`
	Count      int `json:"count"`
	Intent     int `json:"intent"`
}

type Order struct {
	ID         int    `json:"id"`
	CustomerId int    `json:"customerId"`
	ProductId  int    `json:"productId"`
	Count      int    `json:"count"`
	Status     string `json:"status"`
}

type Payment struct {
	ID        int    `json:"id"`
	OrderId   int    `json:"orderId"`
	ProductId int    `json:"productId"`
	TotalCost int    `json:"totalCost"`
	Status    string `json:"status"`
}

type Product struct {
	ID   int    `json:"id"`
	Cost int    `json:"cost"`
	Name string `json:"name"`
}

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PlainIntent struct {
	IntentNumber int `json:"intentNumber"`
}

type Intent struct {
	ID           int `json:"id"`
	IntentNumber int `json:"intentNumber"`
	CustomerId   int `json:"customerId"`
	ProductId    int `json:"productId"`
	OrderId      int `json:"orderId"`
}

// Includes provided order info and all context fetched from DB at start
type OrderWithContext struct {
	CustomerId   int    `json:"customerId"`
	ProductId    int    `json:"productId"`
	Count        int    `json:"count"`
	Status       string `json:"status"`
	IntentNumber int    `json:"intentNumber"`
	ProductCost  int    `json:"productCost"`
}
