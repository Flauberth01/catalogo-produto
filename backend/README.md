# Backend - Catálogo de Produtos

API REST em Go com Clean Architecture para gerenciamento de produtos e categorias.

## 🏗️ Arquitetura

- **Clean Architecture**: Separação clara entre camadas (model, repository, service, handler)
- **GORM**: ORM para PostgreSQL
- **Gin**: Framework web para rotas HTTP
- **PostgreSQL**: Banco de dados relacional

## 📁 Estrutura do Projeto

```
backend/
├── cmd/                    # Ponto de entrada da aplicação
├── internal/              # Código interno da aplicação
│   ├── handler/           # Handlers HTTP (endpoints)
│   ├── service/           # Regras de negócio
│   ├── repository/        # Acesso ao banco de dados
│   └── model/             # Modelos e DTOs
├── db/                    # Configuração do banco de dados
├── main.go               # Arquivo principal
├── Dockerfile            # Containerização
└── go.mod               # Dependências Go
```

## 🚀 Como Executar

### Pré-requisitos

- Go 1.21+
- PostgreSQL 15+
- Docker e Docker Compose (opcional)

### Opção 1: Execução Local

1. **Configurar banco de dados PostgreSQL**
   ```sql
   CREATE DATABASE catalogo_produtos;
   ```

2. **Configurar variáveis de ambiente**
   ```bash
   # Criar arquivo .env baseado no .env.example
   cp .env.example .env
   ```

3. **Instalar dependências**
   ```bash
   go mod download
   ```

4. **Executar aplicação**
   ```bash
   go run main.go
   ```

### Opção 2: Docker Compose (Recomendado)

1. **Executar com Docker Compose**
   ```bash
   docker-compose up -d
   ```

2. **Verificar logs**
   ```bash
   docker-compose logs -f backend
   ```

## 📡 Endpoints da API

### Produtos

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/api/products` | Listar produtos (com filtros opcionais) |
| GET | `/api/products/:id` | Buscar produto por ID |
| POST | `/api/products` | Criar novo produto |
| PUT | `/api/products/:id` | Atualizar produto |
| DELETE | `/api/products/:id` | Remover produto |

### Categorias

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/api/categories` | Listar categorias |
| GET | `/api/categories/:id` | Buscar categoria por ID |
| POST | `/api/categories` | Criar nova categoria |
| PUT | `/api/categories/:id` | Atualizar categoria |
| DELETE | `/api/categories/:id` | Remover categoria |

### Health Check

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/health` | Verificar status da API |

## 🔧 Filtros de Produtos

### Query Parameters

- `name`: Filtrar por nome do produto (busca parcial)
- `category`: Filtrar por nome da categoria (busca parcial)

### Exemplos

```bash
# Buscar produtos com "phone" no nome
GET /api/products?name=phone

# Buscar produtos da categoria "Eletrônicos"
GET /api/products?category=Eletrônicos

# Combinar filtros
GET /api/products?name=phone&category=Eletrônicos
```

## 📊 Modelos de Dados

### Product

```json
{
  "id": 1,
  "name": "Smartphone Galaxy S23",
  "image": "https://example.com/image.jpg",
  "price": 2999.99,
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "Eletrônicos"
  },
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Category

```json
{
  "id": 1,
  "name": "Eletrônicos",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## 🌱 Seed de Dados

A aplicação inclui dados de exemplo que são carregados automaticamente:

- **Categorias**: Eletrônicos, Roupas, Livros, Casa e Jardim, Esportes
- **Produtos**: 8 produtos de exemplo com imagens do Unsplash

## 🔒 CORS

A API está configurada para aceitar requisições de qualquer origem (CORS habilitado).

## 🐳 Deploy

### Render

1. Conectar repositório no Render
2. Configurar como "Web Service"
3. Definir variáveis de ambiente:
   - `DB_HOST`
   - `DB_USER`
   - `DB_PASSWORD`
   - `DB_NAME`
   - `DB_PORT`
   - `PORT`

### Railway

1. Conectar repositório no Railway
2. Adicionar serviço PostgreSQL
3. Configurar variáveis de ambiente automaticamente

### Fly.io

1. Instalar Fly CLI
2. Executar `fly launch`
3. Configurar variáveis de ambiente
4. Deploy com `fly deploy`

## 🧪 Testes

```bash
# Executar testes
go test ./...

# Executar testes com coverage
go test -cover ./...
```

## 📝 Logs

A aplicação utiliza logs estruturados para facilitar o debugging:

- Conexão com banco de dados
- Criação de dados via seed
- Inicialização do servidor
- Erros de validação e negócio 