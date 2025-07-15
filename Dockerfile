# Etapa de build
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila o binário dentro do container Linux
RUN go build -o main ./cmd/main.go

# Etapa final com a mesma base do builder (evita erro de GLIBC)
FROM golang:1.24

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]
