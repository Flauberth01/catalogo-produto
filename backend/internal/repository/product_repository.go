package repository

import (
	"catalogo-produtos/backend/internal/model"

	"gorm.io/gorm"
)

// ProductRepository define as operações de acesso a dados para produtos
type ProductRepository interface {
	Create(product *model.Product) error
	GetByID(id uint) (*model.Product, error)
	GetAll(filters *model.ProductFilter) ([]model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
}

// productRepository implementa ProductRepository
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository cria uma nova instância de ProductRepository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Create cria um novo produto
func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

// GetByID busca um produto pelo ID
func (r *productRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAll busca todos os produtos com filtros opcionais
func (r *productRepository) GetAll(filters *model.ProductFilter) ([]model.Product, error) {
	var products []model.Product
	query := r.db.Preload("Category")

	if filters != nil {
		if filters.Name != "" {
			query = query.Where("name ILIKE ?", "%"+filters.Name+"%")
		}
		if filters.Category != "" {
			query = query.Joins("JOIN categories ON products.category_id = categories.id").
				Where("categories.name ILIKE ?", "%"+filters.Category+"%")
		}
	}

	err := query.Find(&products).Error
	return products, err
}

// Update atualiza um produto existente
func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

// Delete remove um produto pelo ID (soft delete)
func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}
