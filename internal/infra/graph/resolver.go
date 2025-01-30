package graph

import "github.com/devfullcycle/20-CleanArch/internal/usecase"

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
}
