# Catálogo de Produtos

Sistema completo para cadastro, consulta e gerenciamento de produtos e categorias, com backend em Go (Clean Architecture) e frontend em React com TypeScript.

## ✨ Funcionalidades

### 🛍️ Gerenciamento de Produtos
- Cadastro, edição e remoção de produtos
- Upload e exibição de imagens de produtos
- Busca em tempo real por nome e descrição
- Filtros por categorias
- Paginação e ordenação

### 🛒 Sistema de Carrinho
- Adicionar/remover produtos do carrinho
- Atualizar quantidades
- Persistência local no navegador
- Cálculo automático de totais
- Interface intuitiva para gerenciamento

### 📱 Interface Responsiva
- Design adaptativo para mobile e desktop
- Sidebar colapsível em dispositivos móveis
- Componentes acessíveis e modernos
- Loading states e feedback visual
- Navegação intuitiva

### 🔧 Funcionalidades Técnicas
- Integração com banco de dados PostgreSQL
- API documentada com Swagger
- Cache inteligente com React Query
- Arquitetura limpa e escalável
- Injeção de dependências

## 🏗️ Tecnologias Utilizadas

### Backend
- **Go** - Linguagem principal
- **Gin** - Framework web
- **GORM** - ORM para banco de dados
- **Swagger** - Documentação da API
- **Clean Architecture** - Padrão arquitetural

### Frontend
- **React 18** - Biblioteca de UI
- **TypeScript** - Tipagem estática
- **Vite** - Build tool e dev server
- **TailwindCSS** - Framework CSS
- **shadcn/ui** - Componentes UI
- **React Query** - Gerenciamento de estado servidor
- **React Router** - Roteamento
- **Axios** - Cliente HTTP
- **React Hook Form** - Gerenciamento de formulários
- **Zod** - Validação de schemas

### Infraestrutura
- **PostgreSQL** - Banco de dados
- **Docker** - Containerização
- **Docker Compose** - Orquestração

## 📦 Estrutura do Projeto

```
catalogo-produtos/
├── backend/                    # API REST em Go
│   ├── cmd/api/main.go        # Ponto de entrada da aplicação
│   ├── internal/              # Código interno da aplicação
│   │   ├── app/              # Configuração da aplicação
│   │   ├── config/           # Configurações
│   │   ├── domain/           # Regras de negócio
│   │   ├── handler/          # Handlers HTTP
│   │   ├── infrastructure/   # Implementações externas
│   │   ├── model/            # Modelos de dados
│   │   ├── presentation/     # Camada de apresentação
│   │   ├── repository/       # Repositórios
│   │   └── service/          # Serviços
│   ├── db/                   # Scripts de banco de dados
│   ├── docs/                 # Documentação Swagger
│   └── Dockerfile
└── frontend/                 # Aplicação web em React
    ├── src/
    │   ├── components/       # Componentes React
    │   ├── context/          # Contextos (CartContext)
    │   ├── domain/           # Entidades e regras de negócio
    │   ├── hooks/            # Hooks customizados
    │   ├── infrastructure/   # Implementações externas
    │   ├── pages/            # Páginas da aplicação
    │   ├── presentation/     # Hooks de apresentação
    │   └── lib/              # Utilitários
    ├── public/               # Arquivos estáticos
    └── Dockerfile
```

## 🚀 Como rodar o projeto

### Usando Docker Compose (Recomendado)

1. **Clone o repositório:**
   ```sh
   git clone https://github.com/Flauberth01/catalogo-produtos.git
   cd catalogo-produtos
   ```

2. **Configure as variáveis de ambiente:**
   ```sh
   cp backend/env.example backend/.env
   # Edite o arquivo .env conforme necessário
   ```

3. **Suba os containers:**
   ```sh
   docker-compose up --build
   ```

4. **Acesse as aplicações:**
   - **Frontend:** [http://localhost:5173](http://localhost:5173)
   - **API:** [http://localhost:8080/api](http://localhost:8080/api)
   - **Swagger:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Rodando localmente (sem Docker)

#### Pré-requisitos
- Go 1.21+
- Node.js 18+
- PostgreSQL 14+

#### Backend
```sh
cd backend
cp env.example .env
# Configure as variáveis no .env
go mod download
go run cmd/api/main.go
```

#### Frontend
```sh
cd frontend
npm install
npm run dev
```

## 🛒 Funcionalidades do Carrinho

O sistema inclui um carrinho de compras completo com as seguintes funcionalidades:

- **Adicionar produtos:** Clique no botão "Adicionar" em qualquer produto
- **Gerenciar quantidades:** Use os controles +/- no carrinho
- **Remover produtos:** Clique no ícone de lixeira
- **Persistência:** Dados salvos automaticamente no navegador
- **Cálculo automático:** Totais atualizados em tempo real
- **Interface responsiva:** Funciona perfeitamente em mobile e desktop

## 🔍 Busca e Filtros

- **Busca em tempo real:** Digite para buscar por nome ou descrição
- **Filtro por categoria:** Use a sidebar para filtrar produtos
- **Contadores:** Veja quantos produtos foram encontrados
- **Performance:** Filtros aplicados localmente para melhor velocidade

## 📚 Documentação da API

Acesse a documentação interativa do Swagger em:  
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Endpoints Principais

| Método | Endpoint                | Descrição                  |
|--------|------------------------|----------------------------|
| GET    | /api/products          | Lista produtos             |
| GET    | /api/products/:id      | Detalha produto            |
| POST   | /api/products          | Cria produto               |
| PUT    | /api/products/:id      | Atualiza produto           |
| DELETE | /api/products/:id      | Remove produto             |
| GET    | /api/categories        | Lista categorias           |
| POST   | /api/categories        | Cria categoria             |
| PUT    | /api/categories/:id    | Atualiza categoria         |
| DELETE | /api/categories/:id    | Remove categoria           |

## 🏛️ Arquitetura

### Backend (Clean Architecture)
- **Domain:** Entidades e regras de negócio
- **Use Cases:** Casos de uso da aplicação
- **Repository:** Abstração para acesso a dados
- **Infrastructure:** Implementações concretas
- **Presentation:** Controllers e handlers

### Frontend (Clean Architecture)
- **Domain:** Entidades e interfaces
- **Use Cases:** Lógica de negócio
- **Infrastructure:** APIs e repositórios
- **Presentation:** Componentes e hooks
- **DI Container:** Injeção de dependências

## 🛠️ Scripts Disponíveis

### Frontend
```sh
npm run dev          # Inicia servidor de desenvolvimento
npm run build        # Build para produção
npm run build:dev    # Build para desenvolvimento
npm run lint         # Executa ESLint
npm run preview      # Preview do build
```

### Backend
```sh
go run cmd/api/main.go    # Executa a aplicação
go test ./...             # Executa testes
go mod tidy               # Limpa dependências
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

**Flauberth** - [GitHub](https://github.com/Flauberth01)

---

⭐ Se este projeto te ajudou, considere dar uma estrela no repositório!