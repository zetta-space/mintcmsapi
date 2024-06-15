package main

type Product struct {
	ID             uint64 `json:"id" db:"id"`
	SKUNumber      string `json:"sku_number" db:"sku_number"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Price          int    `json:"price" db:"price"`
	Quantity       int    `json:"quantity" db:"quantity"`
	IsSoldOut      bool   `json:"is_sold_out" db:"is_sold_out"`
	IsArchived     bool   `json:"is_archived" db:"is_archived"`
	IsDiscounted   bool   `json:"is_discounted" db:"is_discounted"`
	DiscountAmount int    `json:"discount_amount" db:"discount_amount"`
	DeliveryDays   int    `json:"delivery_days" db:"delivery_days"`
	Remarks        string `json:"remarks" db:"remarks"`
}
