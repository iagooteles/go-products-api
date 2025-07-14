# 🛒 Products REST API - Go (Golang)

Esta é uma API REST desenvolvida em Go com o framework [Gin](https://github.com/gin-gonic/gin), seguindo a arquitetura em camadas (Controller → Usecase → Repository). A aplicação permite consultar produtos cadastrados no banco de dados.

---

## 🚀 Endpoints

| Método | Rota         | Descrição               |
|--------|--------------|-------------------------|
| GET    | `/`          | Rota Teste              |
| GET    | `/products`  | Lista todos os produtos |

---

## 🧱 Estrutura de Pastas

├── cmd/ # Ponto de entrada (main.go)  
├── controller/ # Lida com as requisições HTTP  
├── usecase/   
├── repository/ # Acesso ao banco de dados  
├── db/ # Conexão com o banco  
├── model/ # Estruturas dos dados    
├── go.mod / go.sum # Gerenciamento de dependências  

---

## ⚙️ Tecnologias Utilizadas
- [Go (Golang)](https://go.dev/)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [PostgreSQL](https://www.postgresql.org/)
- Arquitetura limpa: separação por responsabilidades


### Pré-requisitos:

- [Go](https://go.dev/doc/install)
- [Docker + Docker Compose](https://docs.docker.com/get-docker/)


### 1. Clone o projeto:
```bash
git clone https://github.com/iagooteles/go-products-api.git
cd go-products-api
```

### 2. Inicie o banco de dados com Docker:
```bash
docker compose up -d
```

3. Rode a aplicação:

```bash
go run cmd/main.go
```

A API estará disponível em: http://localhost:8000