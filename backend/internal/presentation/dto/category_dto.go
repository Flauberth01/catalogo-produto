package dto

// CategoryCreateRequest representa os dados para criar uma categoria
type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

// CategoryUpdateRequest representa os dados para atualizar uma categoria
type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

// CategoriesResponse representa a resposta de uma lista de categorias
type CategoriesResponse struct {
	Data  []CategoryResponse `json:"data"`
	Total int                `json:"total"`
}

// SingleCategoryResponse representa a resposta de uma categoria única
type SingleCategoryResponse struct {
	Data CategoryResponse `json:"data"`
}
