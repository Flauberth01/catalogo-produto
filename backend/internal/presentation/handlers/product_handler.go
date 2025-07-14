package handlers

import (
	"catalogo-produtos/backend/internal/domain/entities"
	"catalogo-produtos/backend/internal/domain/repositories"
	"catalogo-produtos/backend/internal/domain/usecases"
	"catalogo-produtos/backend/internal/presentation/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ProductHandler gerencia os endpoints HTTP para produtos
type ProductHandler struct {
	productUseCase usecases.ProductUseCase
}

// NewProductHandler cria uma nova instância de ProductHandler
func NewProductHandler(productUseCase usecases.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

// GetProducts retorna todos os produtos com filtros opcionais
// @Summary Listar produtos
// @Description Retorna todos os produtos com filtros opcionais por nome e categoria
// @Tags products
// @Accept json
// @Produce json
// @Param name query string false "Filtrar por nome do produto"
// @Param category query string false "Filtrar por nome da categoria"
// @Success 200 {object} dto.ProductsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var filterReq dto.ProductFilterRequest
	if err := c.ShouldBindQuery(&filterReq); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Parâmetros de filtro inválidos"})
		return
	}

	// Converter DTO para domínio
	filters := &repositories.ProductFilter{
		Name:     filterReq.Name,
		Category: filterReq.Category,
	}

	products, err := h.productUseCase.GetProducts(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Erro ao buscar produtos"})
		return
	}

	// Converter entidades para DTOs de resposta
	productResponses := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = h.mapToProductResponse(product)
	}

	c.JSON(http.StatusOK, dto.ProductsResponse{
		Data:  productResponses,
		Total: len(productResponses),
	})
}

// GetProduct retorna um produto específico pelo ID
// @Summary Buscar produto por ID
// @Description Retorna um produto específico pelo ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} dto.SingleProductResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "ID inválido"})
		return
	}

	product, err := h.productUseCase.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, dto.SingleProductResponse{
		Data: h.mapToProductResponse(*product),
	})
}

// CreateProduct cria um novo produto
// @Summary Criar produto
// @Description Cria um novo produto
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.ProductCreateRequest true "Dados do produto"
// @Success 201 {object} dto.SingleProductResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Dados inválidos"})
		return
	}

	product, err := h.productUseCase.CreateProduct(req.Name, req.Image, req.Description, req.Price, req.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.SingleProductResponse{
		Data: h.mapToProductResponse(*product),
	})
}

// UpdateProduct atualiza um produto existente
// @Summary Atualizar produto
// @Description Atualiza um produto existente
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body dto.ProductUpdateRequest true "Dados do produto"
// @Success 200 {object} dto.SingleProductResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "ID inválido"})
		return
	}

	var req dto.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Dados inválidos"})
		return
	}

	product, err := h.productUseCase.UpdateProduct(uint(id), req.Name, req.Image, req.Description, req.Price, req.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SingleProductResponse{
		Data: h.mapToProductResponse(*product),
	})
}

// DeleteProduct remove um produto
// @Summary Deletar produto
// @Description Remove um produto
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "ID inválido"})
		return
	}

	err = h.productUseCase.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Produto removido com sucesso"})
}

// mapToProductResponse converte entidade para DTO de resposta
func (h *ProductHandler) mapToProductResponse(product entities.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Image:      product.Image,
		Price:      product.Price,
		CategoryID: product.CategoryID,
		Category: dto.CategoryResponse{
			ID:        product.Category.ID,
			Name:      product.Category.Name,
			CreatedAt: product.Category.CreatedAt.Format(time.RFC3339),
			UpdatedAt: product.Category.UpdatedAt.Format(time.RFC3339),
		},
		Description: product.Description,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}
}
