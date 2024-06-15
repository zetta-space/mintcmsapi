package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database interface {
	CreateProduct(*Product) error
	UpdateProduct(*Product) error
	DeleteProduct(int) error
	ListProducts() error
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
	return nil
}

func (p *PostgresDBHandler) UpdateProduct(product *Product) error {
	return nil
}

func (p *PostgresDBHandler) DeleteProduct(id int) error {
	return nil
}

func (p *PostgresDBHandler) ListProducts() error {
	return nil
}
