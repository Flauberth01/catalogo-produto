package handlers

import (
	"catalogo-produtos/backend/internal/domain/entities"
	"catalogo-produtos/backend/internal/domain/usecases"
	"catalogo-produtos/backend/internal/presentation/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CategoryHandler gerencia os endpoints HTTP para categorias
type CategoryHandler struct {
	categoryUseCase usecases.CategoryUseCase
}

// NewCategoryHandler cria uma nova instância de CategoryHandler
func NewCategoryHandler(categoryUseCase usecases.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		categoryUseCase: categoryUseCase,
	}
}

// GetCategories retorna todas as categorias
// @Summary Listar categorias
// @Description Retorna todas as categorias
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} dto.CategoriesResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryUseCase.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Erro ao buscar categorias"})
		return
	}

	// Converter entidades para DTOs de resposta
	categoryResponses := make([]dto.CategoryResponse, len(categories))
	for i, category := range categories {
		categoryResponses[i] = h.mapToCategoryResponse(category)
	}

	c.JSON(http.StatusOK, dto.CategoriesResponse{
		Data:  categoryResponses,
		Total: len(categoryResponses),
	})
}

// GetCategory retorna uma categoria específica pelo ID
// @Summary Buscar categoria por ID
// @Description Retorna uma categoria específica pelo ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID da categoria"
// @Success 200 {object} dto.SingleCategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "ID inválido"})
		return
	}

	category, err := h.categoryUseCase.GetCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Categoria não encontrada"})
		return
	}

	c.JSON(http.StatusOK, dto.SingleCategoryResponse{
		Data: h.mapToCategoryResponse(*category),
	})
}

// CreateCategory cria uma nova categoria
// @Summary Criar categoria
// @Description Cria uma nova categoria
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CategoryCreateRequest true "Dados da categoria"
// @Success 201 {object} dto.SingleCategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Dados inválidos"})
		return
	}

	category, err := h.categoryUseCase.CreateCategory(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.SingleCategoryResponse{
		Data: h.mapToCategoryResponse(*category),
	})
}

// UpdateCategory atualiza uma categoria existente
// @Summary Atualizar categoria
// @Description Atualiza uma categoria existente
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID da categoria"
// @Param category body dto.CategoryUpdateRequest true "Dados da categoria"
// @Success 200 {object} dto.SingleCategoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "ID inválido"})
		return
	}

	var req dto.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Dados inválidos"})
		return
	}

	category, err := h.categoryUseCase.UpdateCategory(uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SingleCategoryResponse{
		Data: h.mapToCategoryResponse(*category),
	})
}

// DeleteCategory remove uma categoria
// @Summary Deletar categoria
// @Description Remove uma categoria
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID da categoria"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "ID inválido"})
		return
	}

	err = h.categoryUseCase.DeleteCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Categoria removida com sucesso"})
}

// mapToCategoryResponse converte entidade para DTO de resposta
func (h *CategoryHandler) mapToCategoryResponse(category entities.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
		UpdatedAt: category.UpdatedAt.Format(time.RFC3339),
	}
}
