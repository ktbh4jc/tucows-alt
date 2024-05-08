package main

// Currently doing 1 writer for all data types. Potential improvement could be splitting
// out into customer/product/order/payment writers.
// Using interface pattern so that we could migrate database systems with relativly little effort
// This allows my consumer to be DB agnostic
type Writer interface {
	CreateOrder(*Order) error
	DeleteOrder(int) error
	UpdateOrder(*Order) error

	CreatePayment(*Payment) error
	DeletePayment(int) error
	UpdatePayment(*Payment) error

	CreateProduct(*Product) error
	DeleteProduct(int) error
	UpdateProduct(*Product) error

	CreateCustomer(*Customer) error
	DeleteCustomer(int) error
	UpdateCustomer(*Customer) error

	//Note: I include a full reset because I think it could be useful for testing
	// I do not recomend doing this IRL.
	FullReset() error

	Init() error
}

type WriterService struct {
	Writer Writer
}

func NewWriterService(writer Writer) *WriterService {
	return &WriterService{
		Writer: writer,
	}
}
