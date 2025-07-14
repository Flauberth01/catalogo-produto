package entities

import "time"

// Product representa a entidade de domínio de um produto
type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"category"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Category representa a entidade de domínio de uma categoria
type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
