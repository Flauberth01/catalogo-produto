package model

import (
	"time"

	"gorm.io/gorm"
)

// Category representa uma categoria de produtos
type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null;size:255;unique"`
	Products  []Product      `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// CategoryCreateRequest representa os dados para criar uma categoria
type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

// CategoryUpdateRequest representa os dados para atualizar uma categoria
type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
