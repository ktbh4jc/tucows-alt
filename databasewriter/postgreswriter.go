package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresWriter struct {
	db *sql.DB
}

func NewPostgresWriter() (*PostgresWriter, error) {
	connStr := "user=postgres dbname=postgres password=SuperSecurePassword2 sslmode=disable host=db"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresWriter{
		db: db,
	}, nil
}

func (pw *PostgresWriter) CreateOrder(order *Order) error {
	query := `
	insert into orders (customerId, productId, count, status)
	values ($1, $2, $3, $4)`
	resp, err := pw.db.Query(query, order.CustomerId, order.ProductId, order.Count, order.Status)

	if err != nil {
		return err
	}
	log.Printf("%+v\n", resp)

	return nil
}

// As intents are a form of record, I only want to be able to create new ones and read existing ones.
func (pw *PostgresWriter) CreateIntent(*Intent) error { return nil }

// Note: I include a full reset because I think it could be useful for testing
// I do not recommend doing this IRL.
func (pw *PostgresWriter) FullReset() error { return nil }

func (pw *PostgresWriter) Init() error {
	err := pw.CreateOrderTable()
	if err != nil {
		return err
	}
	err = pw.CreatePaymentTable()
	if err != nil {
		return err
	}
	err = pw.CreateProductTable()
	if err != nil {
		return err
	}
	err = pw.CreateCustomerTable()
	if err != nil {
		return err
	}
	err = pw.CreateIntentTable()
	return err
}

func (pw *PostgresWriter) CreateOrderTable() error {
	query := `create table if not exists orders (
		id serial primary key,
		customerId serial,
		productId serial,
		count serial,
		status varchar(20)
	)`

	_, err := pw.db.Exec(query)
	return err
}
func (pw *PostgresWriter) CreatePaymentTable() error {
	query := `create table if not exists payments (
		id serial primary key,
		orderId serial,
		productId serial,
		totalCost serial,
		status varchar(20)
	)`

	_, err := pw.db.Exec(query)
	return err
}
func (pw *PostgresWriter) CreateProductTable() error {
	query := `create table if not exists products (
		id serial primary key,
		cost serial,
		name varchar(100)
	)`

	_, err := pw.db.Exec(query)
	return err
}
func (pw *PostgresWriter) CreateCustomerTable() error {
	query := `create table if not exists customers (
		id serial primary key,
		name varchar(100)
	)`

	_, err := pw.db.Exec(query)
	return err
}
func (pw *PostgresWriter) CreateIntentTable() error {
	query := `create table if not exists intents (
		id serial primary key,
		intentNumber serial,
		customerId serial,
		productId serial,
		orderId serial
	)`

	_, err := pw.db.Exec(query)
	return err
}

// Below are the functions I didn't get to for the take-home. I might come back and fill them out
// to have the project completed for my own sake.
func (pw *PostgresWriter) DeleteOrder(int) error    { return nil }
func (pw *PostgresWriter) UpdateOrder(*Order) error { return nil }

func (pw *PostgresWriter) CreatePayment(*Payment) error { return nil }
func (pw *PostgresWriter) DeletePayment(int) error      { return nil }
func (pw *PostgresWriter) UpdatePayment(*Payment) error { return nil }

func (pw *PostgresWriter) CreateProduct(*Product) error { return nil }
func (pw *PostgresWriter) DeleteProduct(int) error      { return nil }
func (pw *PostgresWriter) UpdateProduct(*Product) error { return nil }

func (pw *PostgresWriter) CreateCustomer(*Customer) error { return nil }
func (pw *PostgresWriter) DeleteCustomer(int) error       { return nil }
func (pw *PostgresWriter) UpdateCustomer(*Customer) error { return nil }
