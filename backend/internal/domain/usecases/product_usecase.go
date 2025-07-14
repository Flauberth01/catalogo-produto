package usecases

import (
	"catalogo-produtos/backend/internal/domain/entities"
	"catalogo-produtos/backend/internal/domain/repositories"
	"errors"
)

// ProductUseCase define os casos de uso para produtos
type ProductUseCase interface {
	CreateProduct(name, image, description string, price float64, categoryID uint) (*entities.Product, error)
	GetProduct(id uint) (*entities.Product, error)
	GetProducts(filters *repositories.ProductFilter) ([]entities.Product, error)
	UpdateProduct(id uint, name, image, description string, price float64, categoryID uint) (*entities.Product, error)
	DeleteProduct(id uint) error
}

// productUseCase implementa ProductUseCase
type productUseCase struct {
	productRepo  repositories.ProductRepository
	categoryRepo repositories.CategoryRepository
}

// NewProductUseCase cria uma nova instância de ProductUseCase
func NewProductUseCase(productRepo repositories.ProductRepository, categoryRepo repositories.CategoryRepository) ProductUseCase {
	return &productUseCase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// CreateProduct cria um novo produto
func (uc *productUseCase) CreateProduct(name, image, description string, price float64, categoryID uint) (*entities.Product, error) {
	// Validar se a categoria existe
	_, err := uc.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, errors.New("categoria não encontrada")
	}

	// Validar preço
	if price <= 0 {
		return nil, errors.New("preço deve ser maior que zero")
	}

	// Validar nome
	if name == "" {
		return nil, errors.New("nome é obrigatório")
	}

	product := &entities.Product{
		Name:        name,
		Image:       image,
		Price:       price,
		CategoryID:  categoryID,
		Description: description,
	}

	err = uc.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetProduct busca um produto por ID
func (uc *productUseCase) GetProduct(id uint) (*entities.Product, error) {
	product, err := uc.productRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("produto não encontrado")
	}
	return product, nil
}

// GetProducts busca produtos com filtros
func (uc *productUseCase) GetProducts(filters *repositories.ProductFilter) ([]entities.Product, error) {
	products, err := uc.productRepo.GetAll(filters)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct atualiza um produto
func (uc *productUseCase) UpdateProduct(id uint, name, image, description string, price float64, categoryID uint) (*entities.Product, error) {
	// Buscar produto existente
	product, err := uc.productRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("produto não encontrado")
	}

	// Validar se a categoria existe
	_, err = uc.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, errors.New("categoria não encontrada")
	}

	// Validar preço
	if price <= 0 {
		return nil, errors.New("preço deve ser maior que zero")
	}

	// Validar nome
	if name == "" {
		return nil, errors.New("nome é obrigatório")
	}

	// Atualizar campos
	product.Name = name
	product.Image = image
	product.Price = price
	product.CategoryID = categoryID
	product.Description = description

	err = uc.productRepo.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// DeleteProduct remove um produto
func (uc *productUseCase) DeleteProduct(id uint) error {
	// Verificar se o produto existe
	_, err := uc.productRepo.GetByID(id)
	if err != nil {
		return errors.New("produto não encontrado")
	}

	return uc.productRepo.Delete(id)
}
