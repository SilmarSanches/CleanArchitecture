package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *GetOrdersUseCase) Execute() ([]OrderInputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var orderDTOs []OrderInputDTO
	for _, order := range orders {
		orderDTO := OrderInputDTO{
			ID:    order.ID,
			Price: order.Price,
			Tax:   order.Tax,
		}
		orderDTOs = append(orderDTOs, orderDTO)
	}

	return orderDTOs, nil
}
