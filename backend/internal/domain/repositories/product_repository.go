package repositories

import "catalogo-produtos/backend/internal/domain/entities"

// ProductRepository define as operações de persistência para produtos
type ProductRepository interface {
	Create(product *entities.Product) error
	GetByID(id uint) (*entities.Product, error)
	GetAll(filters *ProductFilter) ([]entities.Product, error)
	Update(product *entities.Product) error
	Delete(id uint) error
}

// ProductFilter define os filtros para busca de produtos
type ProductFilter struct {
	Name     string
	Category string
}
