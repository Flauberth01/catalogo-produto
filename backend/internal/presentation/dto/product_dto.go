package dto

// ProductCreateRequest representa os dados para criar um produto
type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	CategoryID  uint    `json:"category_id" binding:"required"`
	Description string  `json:"description"`
}

// ProductUpdateRequest representa os dados para atualizar um produto
type ProductUpdateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	CategoryID  uint    `json:"category_id" binding:"required"`
	Description string  `json:"description"`
}

// ProductFilterRequest representa os filtros para busca de produtos
type ProductFilterRequest struct {
	Name     string `form:"name"`
	Category string `form:"category"`
}

// ProductResponse representa a resposta de um produto
type ProductResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Image       string           `json:"image"`
	Price       float64          `json:"price"`
	CategoryID  uint             `json:"category_id"`
	Category    CategoryResponse `json:"category"`
	Description string           `json:"description"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
}

// CategoryResponse representa a resposta de uma categoria
type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ProductsResponse representa a resposta de uma lista de produtos
type ProductsResponse struct {
	Data  []ProductResponse `json:"data"`
	Total int               `json:"total"`
}

// ProductResponse representa a resposta de um produto único
type SingleProductResponse struct {
	Data ProductResponse `json:"data"`
}

// MessageResponse representa uma resposta de mensagem
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse representa uma resposta de erro
type ErrorResponse struct {
	Error string `json:"error"`
}
