package handler

import (
	"catalogo-produtos/backend/internal/model"
	"catalogo-produtos/backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductHandler gerencia os endpoints HTTP para produtos
type ProductHandler struct {
	productService service.ProductService
}

// NewProductHandler cria uma nova instância de ProductHandler
func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
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
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var filters model.ProductFilter
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros de filtro inválidos"})
		return
	}

	products, err := h.productService.GetAllProducts(&filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"total": len(products),
	})
}

// GetProduct retorna um produto específico pelo ID
// @Summary Buscar produto por ID
// @Description Retorna um produto específico pelo ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// CreateProduct cria um novo produto
// @Summary Criar produto
// @Description Cria um novo produto
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.ProductCreateRequest true "Dados do produto"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req model.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	product, err := h.productService.CreateProduct(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

// UpdateProduct atualiza um produto existente
// @Summary Atualizar produto
// @Description Atualiza um produto existente
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body model.ProductUpdateRequest true "Dados do produto"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req model.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	product, err := h.productService.UpdateProduct(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct remove um produto
// @Summary Deletar produto
// @Description Remove um produto
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.productService.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido com sucesso"})
}
