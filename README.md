# CleanArchitecture

Projeto de conclusão de pós-graduação (Desafio 3)
Este projeto implementa as consultas de orndens em rest, graphql e gRPC.

## Indice
1. [Docker-Compose](#docker-compose)
2. [Migrations](#migrations)
3. [Rodar a aplicação](#run-server)
4. [Entraga-Rest](#entrega-rest)
5. [Entrga-Graphql](#entrega-graphql)
   + 5.1 [Create](#create)
   + 5.2 [GetAll](#getall)
   + 5.3 [GetByID](#getbyid) 
6. [Comandos gRPC](#comandos-grpc)
   + 6.1 [Problemas com Evans](#problemas-evans)
   + 6.2 [Create](#Create-grpc)
   + 6.3 [GetAll](#GetAll-grpc)
   + 6.4 [GetByID](#GetByID-grpc)
7. [Portas](#portas)

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

## Run-Server
Para iniciar a plicação rode o comando abaixo:

```bash
go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
```

## Entrega-rest

Para a entraga rest foi incluído no projeto um arquivo `get_orders.http` que contém os comandos para realizar as consultas.
Os comandos são dois: um que retorna todas as ordens e outro que consulta por id.

O servidor `rest` roda na porta 8000.

## Entrega-graphql

O servidor grapql roda na porta 8080, onde será exibido o play-ground para execução dos testes

#### Create
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
```

#### GetAll
```bash
query GetAll {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

#### GetByID
```bash
query GetByID {
  order(id: "a") {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## Comandos-grpc

### Problemas Evans
Não consegui rodar o Evans no meu conputador com sucesso, então busquei alternativas.
Consegui tanto com postman quanto com o grpcurl.
Abaixo os comandos para rodar o gRPC com o grpcurl.

#### Create-grpc

```bash
grpcurl -plaintext -d '{"id": "a", "price": 100.0, "tax": 10.0}' localhost:50051 pb.OrderService/CreateOrder
```

#### GetAll-grpc

```bash
grpcurl -plaintext -d '{"id": "a"}' localhost:50051 pb.OrderService/GetOrder
```

#### GetByID-grpc

```bash 
grpcurl -plaintext -d '{}' localhost:50051 pb.OrderService/GetAllOrders  
```

## Portas

As portas utilizadas são:
- 8000: Rest
- 8080: Graphql
- 50051: gRPC