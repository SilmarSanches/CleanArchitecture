# CleanArchitecture

Projeto de conclusão de pós-graduação (Desafio 3)
Este projeto implementa as consultas de orndens em rest, graphql e gRPC.

## Indice
1. [Docker-Compose](#1-docker-compose)
2. [Migrations](#2-migrations)
3. [Rodar a aplicação](#3-rodar-a-aplicação)


## Docker-Compose
Antes de iniciar a aplicação, é necessário subir o banco de dados e o rabbitmq.
Para isso, execute o comando abaixo:

```bash
docker-compose up -d
```
## Migrations
Após subir os recursos acima, aplique as migrations com o comando abaixo:

```bash
make migrate
```

## Rodar a aplicação
Para iniciar a plicação rode o comando abaixo:

```bash
go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
```





### comandos graphql:
```bash
mutation Create {
  createOrder(input: {
    id: "d"
    Price: 105.0
    Tax: 5.0
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}

query GetAll {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}

query GetByID {
  order(id: "a") {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### comandos gRPC usando o grpcurl:
```bash
grpcurl -plaintext -d '{"id": "a", "price": 100.0, "tax": 10.0}' localhost:50051 pb.OrderService/CreateOrder

grpcurl -plaintext -d '{"id": "1"}' localhost:50051 pb.OrderService/GetOrder

grpcurl -plaintext -d '{}' localhost:50051 pb.OrderService/GetAllOrders

```