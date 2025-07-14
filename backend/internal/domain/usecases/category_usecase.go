package usecases

import (
	"catalogo-produtos/backend/internal/domain/entities"
	"catalogo-produtos/backend/internal/domain/repositories"
	"errors"
)

// CategoryUseCase define os casos de uso para categorias
type CategoryUseCase interface {
	CreateCategory(name string) (*entities.Category, error)
	GetCategory(id uint) (*entities.Category, error)
	GetCategories() ([]entities.Category, error)
	UpdateCategory(id uint, name string) (*entities.Category, error)
	DeleteCategory(id uint) error
}

// categoryUseCase implementa CategoryUseCase
type categoryUseCase struct {
	categoryRepo repositories.CategoryRepository
}

// NewCategoryUseCase cria uma nova instância de CategoryUseCase
func NewCategoryUseCase(categoryRepo repositories.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
	}
}

// CreateCategory cria uma nova categoria
func (uc *categoryUseCase) CreateCategory(name string) (*entities.Category, error) {
	// Validar nome
	if name == "" {
		return nil, errors.New("nome é obrigatório")
	}

	category := &entities.Category{
		Name: name,
	}

	err := uc.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategory busca uma categoria por ID
func (uc *categoryUseCase) GetCategory(id uint) (*entities.Category, error) {
	category, err := uc.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("categoria não encontrada")
	}
	return category, nil
}

// GetCategories busca todas as categorias
func (uc *categoryUseCase) GetCategories() ([]entities.Category, error) {
	categories, err := uc.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// UpdateCategory atualiza uma categoria
func (uc *categoryUseCase) UpdateCategory(id uint, name string) (*entities.Category, error) {
	// Buscar categoria existente
	category, err := uc.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("categoria não encontrada")
	}

	// Validar nome
	if name == "" {
		return nil, errors.New("nome é obrigatório")
	}

	// Atualizar nome
	category.Name = name

	err = uc.categoryRepo.Update(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategory remove uma categoria
func (uc *categoryUseCase) DeleteCategory(id uint) error {
	// Verificar se a categoria existe
	_, err := uc.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("categoria não encontrada")
	}

	return uc.categoryRepo.Delete(id)
}
