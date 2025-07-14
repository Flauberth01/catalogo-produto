package main

import (
	"catalogo-produtos/backend/internal/app"
	"catalogo-produtos/backend/internal/config"
	"log"

	"github.com/joho/godotenv"
)

// @title           Catálogo de Produtos API
// @version         1.0
// @description     API REST para gerenciamento de produtos e categorias
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Carregar configurações
	cfg := config.Load()

	// Criar e inicializar aplicação
	application := app.NewApp(cfg)
	if err := application.Initialize(); err != nil {
		log.Fatal("Erro ao inicializar aplicação:", err)
	}
	defer application.Close()

	// Executar aplicação
	if err := application.Run(); err != nil {
		log.Fatal("Erro ao executar aplicação:", err)
	}
}
