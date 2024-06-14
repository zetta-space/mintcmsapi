package main

type Product struct {
	ID        uint64 `json:"id" db:"id"`
	SKUNumber string `json:"sku_number" db:"sku_number"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
}
