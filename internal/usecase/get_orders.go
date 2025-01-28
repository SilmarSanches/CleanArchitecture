package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderGet        events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderGets events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderGet:        OrderGets,
		EventDispatcher: EventDispatcher,
	}
}

func (c *GetOrdersUseCase) GetAll() ([]OrderInputDTO, error) {
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

	c.OrderGet.SetPayload(orderDTOs)
	c.EventDispatcher.Dispatch(c.OrderGet)

	return orderDTOs, nil
}
