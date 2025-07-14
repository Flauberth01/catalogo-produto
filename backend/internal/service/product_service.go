package service

import (
	"catalogo-produtos/backend/internal/model"
	"catalogo-produtos/backend/internal/repository"
	"errors"
)

// ProductService define as operações de negócio para produtos
type ProductService interface {
	CreateProduct(req *model.ProductCreateRequest) (*model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	GetAllProducts(filters *model.ProductFilter) ([]model.Product, error)
	UpdateProduct(id uint, req *model.ProductUpdateRequest) (*model.Product, error)
	DeleteProduct(id uint) error
}

// productService implementa ProductService
type productService struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

// NewProductService cria uma nova instância de ProductService
func NewProductService(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductService {
	return &productService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// CreateProduct cria um novo produto
func (s *productService) CreateProduct(req *model.ProductCreateRequest) (*model.Product, error) {
	// Verifica se a categoria existe
	_, err := s.categoryRepo.GetByID(req.CategoryID)
	if err != nil {
		return nil, errors.New("categoria não encontrada")
	}

	product := &model.Product{
		Name:       req.Name,
		Image:      req.Image,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	}

	err = s.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	// Retorna o produto com a categoria carregada
	return s.productRepo.GetByID(product.ID)
}

// GetProductByID busca um produto pelo ID
func (s *productService) GetProductByID(id uint) (*model.Product, error) {
	return s.productRepo.GetByID(id)
}

// GetAllProducts busca todos os produtos com filtros
func (s *productService) GetAllProducts(filters *model.ProductFilter) ([]model.Product, error) {
	return s.productRepo.GetAll(filters)
}

// UpdateProduct atualiza um produto existente
func (s *productService) UpdateProduct(id uint, req *model.ProductUpdateRequest) (*model.Product, error) {
	// Verifica se o produto existe
	existingProduct, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("produto não encontrado")
	}

	// Verifica se a categoria existe
	_, err = s.categoryRepo.GetByID(req.CategoryID)
	if err != nil {
		return nil, errors.New("categoria não encontrada")
	}

	// Atualiza os campos
	existingProduct.Name = req.Name
	existingProduct.Image = req.Image
	existingProduct.Price = req.Price
	existingProduct.CategoryID = req.CategoryID

	err = s.productRepo.Update(existingProduct)
	if err != nil {
		return nil, err
	}

	// Retorna o produto atualizado com a categoria carregada
	return s.productRepo.GetByID(id)
}

// DeleteProduct remove um produto
func (s *productService) DeleteProduct(id uint) error {
	// Verifica se o produto existe
	_, err := s.productRepo.GetByID(id)
	if err != nil {
		return errors.New("produto não encontrado")
	}

	return s.productRepo.Delete(id)
}
