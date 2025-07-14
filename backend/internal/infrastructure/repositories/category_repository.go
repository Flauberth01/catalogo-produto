package repositories

import (
	"catalogo-produtos/backend/internal/domain/entities"
	"catalogo-produtos/backend/internal/domain/repositories"
	"catalogo-produtos/backend/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// categoryRepository implementa CategoryRepository
type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository cria uma nova instância de CategoryRepository
func NewCategoryRepository(db *gorm.DB) repositories.CategoryRepository {
	return &categoryRepository{db: db}
}

// Create cria uma nova categoria
func (r *categoryRepository) Create(category *entities.Category) error {
	model := &models.CategoryModel{
		Name: category.Name,
	}

	err := r.db.Create(model).Error
	if err != nil {
		return err
	}

	// Atualizar o ID da categoria criada
	category.ID = model.ID
	category.CreatedAt = model.CreatedAt
	category.UpdatedAt = model.UpdatedAt

	return nil
}

// GetByID busca uma categoria por ID
func (r *categoryRepository) GetByID(id uint) (*entities.Category, error) {
	var model models.CategoryModel
	err := r.db.First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return r.mapToEntity(&model), nil
}

// GetAll busca todas as categorias
func (r *categoryRepository) GetAll() ([]entities.Category, error) {
	var models []models.CategoryModel
	err := r.db.Find(&models).Error
	if err != nil {
		return nil, err
	}

	// Converter para entidades
	categories := make([]entities.Category, len(models))
	for i, model := range models {
		categories[i] = *r.mapToEntity(&model)
	}

	return categories, nil
}

// Update atualiza uma categoria
func (r *categoryRepository) Update(category *entities.Category) error {
	model := &models.CategoryModel{
		ID:   category.ID,
		Name: category.Name,
	}

	err := r.db.Save(model).Error
	if err != nil {
		return err
	}

	// Atualizar timestamps
	category.UpdatedAt = model.UpdatedAt

	return nil
}

// Delete remove uma categoria
func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.CategoryModel{}, id).Error
}

// mapToEntity converte modelo para entidade
func (r *categoryRepository) mapToEntity(model *models.CategoryModel) *entities.Category {
	return &entities.Category{
		ID:        model.ID,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
