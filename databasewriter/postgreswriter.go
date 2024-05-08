package main

import "database/sql"

type PostgresWriter struct {
	db *sql.DB
}

func NewPostgresWriter() (*PostgresWriter, error) {
	connStr := "user=postgres dbname=postgres password=SuperSecurePassword2 sslmode=disable"

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

func (pw *PostgresWriter) CreateOrder(*Order) error
func (pw *PostgresWriter) DeleteOrder(int) error
func (pw *PostgresWriter) UpdateOrder(*Order) error

func (pw *PostgresWriter) CreatePayment(*Payment) error
func (pw *PostgresWriter) DeletePayment(int) error
func (pw *PostgresWriter) UpdatePayment(*Payment) error

func (pw *PostgresWriter) CreateProduct(*Product) error
func (pw *PostgresWriter) DeleteProduct(int) error
func (pw *PostgresWriter) UpdateProduct(*Product) error

func (pw *PostgresWriter) CreateCustomer(*Customer) error
func (pw *PostgresWriter) DeleteCustomer(int) error
func (pw *PostgresWriter) UpdateCustomer(*Customer) error

// As intents are a form of record, I only want to be able to create new ones and read existing ones.
func (pw *PostgresWriter) CreateIntent(*Intent) error

// Note: I include a full reset because I think it could be useful for testing
// I do not recomend doing this IRL.
func (pw *PostgresWriter) FullReset() error

func (pw *PostgresWriter) Init() error
