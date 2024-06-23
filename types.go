package main

import "time"

type Product struct {
	ID             string    `json:"id" db:"id"`
	SKUNumber      string    `json:"sku_number" db:"sku_number"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Price          float32   `json:"price" db:"price"`
	Quantity       int       `json:"quantity" db:"quantity"`
	IsSoldOut      bool      `json:"isSoldOut" db:"is_sold_out"`
	IsArchived     bool      `json:"isArchived" db:"is_archived"`
	IsDiscounted   bool      `json:"isDiscounted" db:"is_discounted"`
	DiscountAmount float32   `json:"discountAmount" db:"discount_amount"`
	DeliveryDays   int       `json:"deliveryDays" db:"delivery_days"`
	Remarks        string    `json:"remarks" db:"remarks"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}

type NewProductRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float32 `json:"price"`
	Quantity     int     `json:"quantity"`
	DeliveryDays int     `json:"deliveryDays"`
}

func NewProduct(name, description string, price float32, quantity int, deliveryDays int) *Product {
	return &Product{
		Name:         name,
		Description:  description,
		Price:        price,
		Quantity:     quantity,
		DeliveryDays: deliveryDays,
	}
}
