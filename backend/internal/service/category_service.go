package service

import (
	"catalogo-produtos/backend/internal/model"
	"catalogo-produtos/backend/internal/repository"
	"errors"
)

// CategoryService define as operações de negócio para categorias
type CategoryService interface {
	CreateCategory(req *model.CategoryCreateRequest) (*model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	GetAllCategories() ([]model.Category, error)
	UpdateCategory(id uint, req *model.CategoryUpdateRequest) (*model.Category, error)
	DeleteCategory(id uint) error
}

// categoryService implementa CategoryService
type categoryService struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryService cria uma nova instância de CategoryService
func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

// CreateCategory cria uma nova categoria
func (s *categoryService) CreateCategory(req *model.CategoryCreateRequest) (*model.Category, error) {
	category := &model.Category{
		Name: req.Name,
	}

	err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}

	return s.categoryRepo.GetByID(category.ID)
}

// GetCategoryByID busca uma categoria pelo ID
func (s *categoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.categoryRepo.GetByID(id)
}

// GetAllCategories busca todas as categorias
func (s *categoryService) GetAllCategories() ([]model.Category, error) {
	return s.categoryRepo.GetAll()
}

// UpdateCategory atualiza uma categoria existente
func (s *categoryService) UpdateCategory(id uint, req *model.CategoryUpdateRequest) (*model.Category, error) {
	// Verifica se a categoria existe
	existingCategory, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("categoria não encontrada")
	}

	// Atualiza o nome
	existingCategory.Name = req.Name

	err = s.categoryRepo.Update(existingCategory)
	if err != nil {
		return nil, err
	}

	return s.categoryRepo.GetByID(id)
}

// DeleteCategory remove uma categoria
func (s *categoryService) DeleteCategory(id uint) error {
	// Verifica se a categoria existe
	_, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("categoria não encontrada")
	}

	return s.categoryRepo.Delete(id)
}
