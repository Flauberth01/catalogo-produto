package handler

import (
	"catalogo-produtos/backend/internal/model"
	"catalogo-produtos/backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CategoryHandler gerencia os endpoints HTTP para categorias
type CategoryHandler struct {
	categoryService service.CategoryService
}

// NewCategoryHandler cria uma nova instância de CategoryHandler
func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// GetCategories retorna todas as categorias
// @Summary Listar categorias
// @Description Retorna todas as categorias
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar categorias"})
		return
	}

	// Adiciona a categoria 'Todos' no início
	virtualTodos := model.Category{
		ID:   0,
		Name: "Todos",
	}
	categoriesWithTodos := append([]model.Category{virtualTodos}, categories...)

	c.JSON(http.StatusOK, gin.H{
		"data":  categoriesWithTodos,
		"total": len(categoriesWithTodos),
	})
}

// GetCategory retorna uma categoria específica pelo ID
// @Summary Buscar categoria por ID
// @Description Retorna uma categoria específica pelo ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID da categoria"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	category, err := h.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// CreateCategory cria uma nova categoria
// @Summary Criar categoria
// @Description Cria uma nova categoria
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.CategoryCreateRequest true "Dados da categoria"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req model.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	category, err := h.categoryService.CreateCategory(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": category})
}

// UpdateCategory atualiza uma categoria existente
// @Summary Atualizar categoria
// @Description Atualiza uma categoria existente
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID da categoria"
// @Param category body model.CategoryUpdateRequest true "Dados da categoria"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req model.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	category, err := h.categoryService.UpdateCategory(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory remove uma categoria
// @Summary Deletar categoria
// @Description Remove uma categoria
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID da categoria"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.categoryService.DeleteCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoria removida com sucesso"})
}
