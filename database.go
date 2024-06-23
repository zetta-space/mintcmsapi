package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database interface {
	CreateProduct(*Product) error
	UpdateProduct(*Product) error
	DeleteProduct(int) error
	ListProducts() ([]*Product, error)
}

type PostgresDBHandler struct {
	db *sql.DB
}

func NewConnection() (*PostgresDBHandler, error) {
	connStr := "user=postgres dbname=postgres password=Aspirine@123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	// Ping the db if error happend
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDBHandler{db}, nil
}

// Implement the interface

func (p *PostgresDBHandler) CreateProduct(product *Product) error {
	query := `insert into product(
	sku_number, 
	name, 
	description,
	price,
	quantity,
	is_sold_out,
	is_archived,
	is_discounted,
	discount_amount,
	delivery_days,
	remarks) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	response, err := p.db.Query(query,
		product.SKUNumber,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.IsSoldOut,
		product.IsArchived,
		product.IsDiscounted,
		product.DiscountAmount,
		product.DeliveryDays, product.Remarks)

	if err != nil {
		return err
	}
	// print the new product
	fmt.Printf("%+v\n", response)
	return nil
}

func (p *PostgresDBHandler) UpdateProduct(product *Product) error {
	return nil
}

func (p *PostgresDBHandler) DeleteProduct(id int) error {
	return nil
}

func (p *PostgresDBHandler) ListProducts() ([]*Product, error) {
	rows, err := p.db.Query("select * from product")
	if err != nil {
		return nil, err
	}

	products := []*Product{}
	for rows.Next() {
		product := new(Product)
		err := rows.Scan(
			&product.ID,
			&product.SKUNumber,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
			&product.IsSoldOut,
			&product.IsArchived,
			&product.IsDiscounted,
			&product.DiscountAmount,
			&product.DeliveryDays,
			&product.Remarks,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// Create tables
func (s *PostgresDBHandler) Init() error {
	return s.CreateProductTable()
}

func (s *PostgresDBHandler) CreateProductTable() error {
	// Run extensions which are important
	s.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	query := `CREATE TABLE IF NOT EXISTS product (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	sku_number VARCHAR(255) NULL, 
	name VARCHAR(255),  
	description VARCHAR(255),  
	price DECIMAL(10,2), 
	quantity INT, 
	is_sold_out BOOLEAN DEFAULT FALSE,
	is_archived BOOLEAN DEFAULT FALSE,
	is_discounted BOOLEAN DEFAULT FALSE,
	discount_amount DECIMAL(10,2) NULL,
	delivery_days INT, 
	remarks VARCHAR(255) NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
	)`

	_, err := s.db.Exec(query)
	return err
}
