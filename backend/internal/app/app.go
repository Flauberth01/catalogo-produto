package app

import (
	"catalogo-produtos/backend/db"
	"catalogo-produtos/backend/docs"
	"catalogo-produtos/backend/internal/config"
	"catalogo-produtos/backend/internal/domain/usecases"
	infraRepos "catalogo-produtos/backend/internal/infrastructure/repositories"
	"catalogo-produtos/backend/internal/presentation/handlers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// App representa a aplicação principal
type App struct {
	config *config.Config
	router *gin.Engine
	db     *db.Database
}

// NewApp cria uma nova instância da aplicação
func NewApp(cfg *config.Config) *App {
	return &App{
		config: cfg,
		router: gin.Default(),
	}
}

// Initialize inicializa a aplicação
func (a *App) Initialize() error {
	// Configurar modo do Gin
	gin.SetMode(a.config.Server.Mode)

	// Configurar banco de dados
	database := db.NewDatabase(&a.config.Database)
	a.db = database

	// Executar seed
	db.Seed(database.DB)

	// Configurar CORS
	a.setupCORS()

	// Configurar rotas
	a.setupRoutes()

	return nil
}

// setupCORS configura o CORS
func (a *App) setupCORS() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	a.router.Use(cors.New(config))
}

// setupRoutes configura as rotas da aplicação
func (a *App) setupRoutes() {
	// Configurar repositórios (Infrastructure Layer)
	productRepo := infraRepos.NewProductRepository(a.db.DB)
	categoryRepo := infraRepos.NewCategoryRepository(a.db.DB)

	// Configurar casos de uso (Domain Layer)
	productUseCase := usecases.NewProductUseCase(productRepo, categoryRepo)
	categoryUseCase := usecases.NewCategoryUseCase(categoryRepo)

	// Configurar handlers (Presentation Layer)
	productHandler := handlers.NewProductHandler(productUseCase)
	categoryHandler := handlers.NewCategoryHandler(categoryUseCase)

	// Rotas da API
	api := a.router.Group("/api")
	{
		// Rotas de produtos
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}

		// Rotas de categorias
		categories := api.Group("/categories")
		{
			categories.GET("", categoryHandler.GetCategories)
			categories.GET("/:id", categoryHandler.GetCategory)
			categories.POST("", categoryHandler.CreateCategory)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}
	}

	// Swagger
	docs.SwaggerInfo.Title = "Catálogo de Produtos API"
	docs.SwaggerInfo.Description = "API REST para gerenciamento de produtos e categorias"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	a.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API está funcionando",
		})
	})
}

// Run inicia o servidor
func (a *App) Run() error {
	log.Printf("Servidor iniciado na porta %s", a.config.Server.Port)
	log.Printf("API disponível em: http://localhost:%s/api", a.config.Server.Port)
	log.Printf("Swagger UI: http://localhost:%s/swagger/index.html", a.config.Server.Port)
	log.Printf("Health check: http://localhost:%s/health", a.config.Server.Port)

	return a.router.Run(":" + a.config.Server.Port)
}

// Close fecha a aplicação
func (a *App) Close() {
	if a.db != nil {
		a.db.Close()
	}
}
