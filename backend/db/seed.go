package db

import (
	"catalogo-produtos/backend/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// Seed popula o banco de dados com dados iniciais
func Seed(db *gorm.DB) {
	log.Println("Iniciando seed do banco de dados...")

	// Criar categorias
	categories := []models.CategoryModel{
		{Name: "Eletrônicos"},
		{Name: "Roupas"},
		{Name: "Livros"},
		{Name: "Casa e Jardim"},
		{Name: "Esportes"},
	}

	for _, category := range categories {
		var existingCategory models.CategoryModel
		if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&category).Error; err != nil {
					log.Printf("Erro ao criar categoria %s: %v", category.Name, err)
				} else {
					log.Printf("Categoria criada: %s", category.Name)
				}
			}
		}
	}

	// Buscar categorias criadas para usar nos produtos
	var eletronicos, roupas, livros, casa, esportes models.CategoryModel
	db.Where("name = ?", "Eletrônicos").First(&eletronicos)
	db.Where("name = ?", "Roupas").First(&roupas)
	db.Where("name = ?", "Livros").First(&livros)
	db.Where("name = ?", "Casa e Jardim").First(&casa)
	db.Where("name = ?", "Esportes").First(&esportes)

	// Criar produtos
	products := []models.ProductModel{
		{
			Name:        "Smartphone Galaxy S23",
			Image:       "https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=400",
			Price:       2999.99,
			CategoryID:  eletronicos.ID,
			Description: "Smartphone com tela de 6.4 polegadas, câmera tripla de 50MP e bateria de 5000mAh",
		},
		{
			Name:        "Notebook Dell Inspiron",
			Image:       "https://images.unsplash.com/photo-1496181133206-80ce9b88a853?w=400",
			Price:       4599.99,
			CategoryID:  eletronicos.ID,
			Description: "Notebook com Intel Core i5, 8GB RAM, SSD 256GB e tela Full HD de 15.6 polegadas",
		},
		{
			Name:        "Camiseta Básica",
			Image:       "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400",
			Price:       49.99,
			CategoryID:  roupas.ID,
			Description: "Camiseta 100% algodão com corte moderno e tecido de alta qualidade",
		},
		{
			Name:        "Calça Jeans",
			Image:       "https://images.unsplash.com/photo-1542272604-787c3835535d?w=400",
			Price:       129.99,
			CategoryID:  roupas.ID,
			Description: "Calça jeans tradicional com corte clássico e acabamentos de qualidade",
		},
		{
			Name:        "O Senhor dos Anéis",
			Image:       "https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=400",
			Price:       89.99,
			CategoryID:  livros.ID,
			Description: "Livro clássico de fantasia escrito por J.R.R. Tolkien, edição especial ilustrada",
		},
		{
			Name:        "Vaso Decorativo",
			Image:       "https://images.unsplash.com/photo-1485955900006-10f4d324d411?w=400",
			Price:       79.99,
			CategoryID:  casa.ID,
			Description: "Vaso decorativo em cerâmica para plantas, ideal para ambientes internos",
		},
		{
			Name:        "Bola de Futebol",
			Image:       "https://images.unsplash.com/photo-1552318965-6e6be7484ada?w=400",
			Price:       89.99,
			CategoryID:  esportes.ID,
			Description: "Bola de futebol oficial tamanho 5, costurada à mão, ideal para jogos e treinos",
		},
		{
			Name:        "Tênis Nike",
			Image:       "https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=400",
			Price:       299.99,
			CategoryID:  esportes.ID,
			Description: "Tênis esportivo Nike com tecnologia Air Max para máximo conforto e performance",
		},
	}

	for _, product := range products {
		var existingProduct models.ProductModel
		if err := db.Where("name = ?", product.Name).First(&existingProduct).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&product).Error; err != nil {
					log.Printf("Erro ao criar produto %s: %v", product.Name, err)
				} else {
					log.Printf("Produto criado: %s", product.Name)
				}
			}
		}
	}

	log.Println("Seed concluído!")
}
