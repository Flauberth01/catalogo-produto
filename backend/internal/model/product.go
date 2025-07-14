package model

import (
	"time"

	"gorm.io/gorm"
)

// Product representa um produto no sistema
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:255"`
	Image       string         `json:"image" gorm:"size:500"`
	Price       float64        `json:"price" gorm:"not null;type:decimal(10,2)"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ProductCreateRequest representa os dados para criar um produto
type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	CategoryID  uint    `json:"category_id" binding:"required"`
	Description string  `json:"description"`
}

// ProductUpdateRequest representa os dados para atualizar um produto
type ProductUpdateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	CategoryID  uint    `json:"category_id" binding:"required"`
	Description string  `json:"description"`
}

// ProductFilter representa os filtros para busca de produtos
type ProductFilter struct {
	Name     string `form:"name"`
	Category string `form:"category"`
}
