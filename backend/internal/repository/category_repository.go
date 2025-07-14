package repository

import (
	"catalogo-produtos/backend/internal/model"

	"gorm.io/gorm"
)

// CategoryRepository define as operações de acesso a dados para categorias
type CategoryRepository interface {
	Create(category *model.Category) error
	GetByID(id uint) (*model.Category, error)
	GetAll() ([]model.Category, error)
	Update(category *model.Category) error
	Delete(id uint) error
}

// categoryRepository implementa CategoryRepository
type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository cria uma nova instância de CategoryRepository
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// Create cria uma nova categoria
func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

// GetByID busca uma categoria pelo ID
func (r *categoryRepository) GetByID(id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetAll busca todas as categorias
func (r *categoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

// Update atualiza uma categoria existente
func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

// Delete remove uma categoria pelo ID (soft delete)
func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}
