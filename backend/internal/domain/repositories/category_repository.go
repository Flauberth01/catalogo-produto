package repositories

import "catalogo-produtos/backend/internal/domain/entities"

// CategoryRepository define as operações de persistência para categorias
type CategoryRepository interface {
	Create(category *entities.Category) error
	GetByID(id uint) (*entities.Category, error)
	GetAll() ([]entities.Category, error)
	Update(category *entities.Category) error
	Delete(id uint) error
}
