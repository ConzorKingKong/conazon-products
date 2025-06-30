package types

import "time"

type Product struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MainImage   string    `json:"mainImage"`
	Category    string    `json:"category"`
	Price       float32   `json:"price"`
	Quantity    int       `json:"quantity,omitempty"`
	Author      string    `json:"author,omitempty"`
}

type ProductsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

type ProductResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}