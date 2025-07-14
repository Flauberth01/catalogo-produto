package models

import (
	"time"

	"gorm.io/gorm"
)

// ProductModel representa o modelo de banco de dados para produtos
type ProductModel struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:255"`
	Image       string         `json:"image" gorm:"size:500"`
	Price       float64        `json:"price" gorm:"not null;type:decimal(10,2)"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	Category    CategoryModel  `json:"category" gorm:"foreignKey:CategoryID"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName especifica o nome da tabela
func (ProductModel) TableName() string {
	return "products"
}

// CategoryModel representa o modelo de banco de dados para categorias
type CategoryModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null;size:255;unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName especifica o nome da tabela
func (CategoryModel) TableName() string {
	return "categories"
}
