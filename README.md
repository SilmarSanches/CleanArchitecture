# CleanArchitecture









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