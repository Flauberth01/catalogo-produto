package repositories

import (
	"catalogo-produtos/backend/internal/domain/entities"
	"catalogo-produtos/backend/internal/domain/repositories"
	"catalogo-produtos/backend/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// productRepository implementa ProductRepository
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository cria uma nova instância de ProductRepository
func NewProductRepository(db *gorm.DB) repositories.ProductRepository {
	return &productRepository{db: db}
}

// Create cria um novo produto
func (r *productRepository) Create(product *entities.Product) error {
	model := &models.ProductModel{
		Name:        product.Name,
		Image:       product.Image,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		Description: product.Description,
	}

	err := r.db.Create(model).Error
	if err != nil {
		return err
	}

	// Atualizar o ID do produto criado
	product.ID = model.ID
	product.CreatedAt = model.CreatedAt
	product.UpdatedAt = model.UpdatedAt

	return nil
}

// GetByID busca um produto por ID
func (r *productRepository) GetByID(id uint) (*entities.Product, error) {
	var model models.ProductModel
	err := r.db.Preload("Category").First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return r.mapToEntity(&model), nil
}

// GetAll busca todos os produtos com filtros
func (r *productRepository) GetAll(filters *repositories.ProductFilter) ([]entities.Product, error) {
	var models []models.ProductModel
	query := r.db.Preload("Category")

	// Aplicar filtros
	if filters != nil {
		if filters.Name != "" {
			query = query.Where("name ILIKE ?", "%"+filters.Name+"%")
		}
		if filters.Category != "" {
			query = query.Joins("JOIN categories ON categories.id = products.category_id").
				Where("categories.name ILIKE ?", "%"+filters.Category+"%")
		}
	}

	err := query.Find(&models).Error
	if err != nil {
		return nil, err
	}

	// Converter para entidades
	products := make([]entities.Product, len(models))
	for i, model := range models {
		products[i] = *r.mapToEntity(&model)
	}

	return products, nil
}

// Update atualiza um produto
func (r *productRepository) Update(product *entities.Product) error {
	model := &models.ProductModel{
		ID:          product.ID,
		Name:        product.Name,
		Image:       product.Image,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		Description: product.Description,
	}

	err := r.db.Save(model).Error
	if err != nil {
		return err
	}

	// Atualizar timestamps
	product.UpdatedAt = model.UpdatedAt

	return nil
}

// Delete remove um produto
func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.ProductModel{}, id).Error
}

// mapToEntity converte modelo para entidade
func (r *productRepository) mapToEntity(model *models.ProductModel) *entities.Product {
	return &entities.Product{
		ID:          model.ID,
		Name:        model.Name,
		Image:       model.Image,
		Price:       model.Price,
		CategoryID:  model.CategoryID,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		Category: entities.Category{
			ID:        model.Category.ID,
			Name:      model.Category.Name,
			CreatedAt: model.Category.CreatedAt,
			UpdatedAt: model.Category.UpdatedAt,
		},
	}
}
