# 🛒 Products REST API - Go (Golang)

Esta é uma API REST desenvolvida em Go com o framework [Gin](https://github.com/gin-gonic/gin), seguindo a arquitetura em camadas (Controller → Usecase → Repository). A aplicação permite consultar produtos cadastrados no banco de dados.

---

## 🚀 Endpoints

| Método | Rota                   | Descrição               |
|--------|------------------------|-------------------------|
| GET    | `/`                    | Rota Teste              |
| GET    | `/products`            | Lista todos os produtos |
| POST   | `/product`             | Adiciona um produto     |
| GET    | `/product/:productId`  | Busca produto por id    |
| PUT    | `/product`             | Rota de Update          |

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
- Docker & Docker Compose
- Arquitetura limpa: separação por responsabilidades


### Pré-requisitos:

- [Go](https://go.dev/doc/install)
- [Docker + Docker Compose](https://docs.docker.com/get-docker/)


### 1. Clone o projeto:
```bash
git clone https://github.com/iagooteles/go-products-api.git
cd go-products-api
```

### 2. Inicie o app e o banco de dados com Docker:
```bash
docker compose up --build
```

3. Acesse a API em:

```bash
http://localhost:8000
```

🔧 Rodando sem Docker (modo local)
Apenas se preferir usar o Go instalado na sua máquina

Suba o banco:
```bash
docker compose up -d go_db
```

Rode a aplicação Go:
```bash
go run cmd/main.go
```

OBS: Em `db/conn.go`, usar localhost quando estiver fazendo os testes na máquina sem docker.
    Com docker usar a conexão go_db.