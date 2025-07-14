package db

import (
	"catalogo-produtos/backend/internal/config"
	"catalogo-produtos/backend/internal/infrastructure/database/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database representa a conexão com o banco de dados
type Database struct {
	DB *gorm.DB
}

// NewDatabase cria uma nova conexão com o banco de dados
func NewDatabase(config *config.DatabaseConfig) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}

	// Auto-migrar tabelas
	err = db.AutoMigrate(
		&models.CategoryModel{},
		&models.ProductModel{},
	)
	if err != nil {
		log.Fatal("Erro ao migrar tabelas:", err)
	}

	log.Println("Banco de dados conectado e migrado com sucesso")

	return &Database{DB: db}
}

// Close fecha a conexão com o banco de dados
func (d *Database) Close() {
	sqlDB, err := d.DB.DB()
	if err != nil {
		log.Println("Erro ao obter conexão SQL:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Println("Erro ao fechar conexão com banco de dados:", err)
	}
}
