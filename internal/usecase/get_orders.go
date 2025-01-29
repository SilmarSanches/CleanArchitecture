package usecase

import (
	"fmt"
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

func (c *GetOrdersUseCase) GetAll() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	if orders == nil {
		return nil, fmt.Errorf("nenhum registro encontrado")
	}

	var orderDTOs []OrderOutputDTO
	for _, order := range orders {
		orderDTO := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		orderDTOs = append(orderDTOs, orderDTO)
	}

	c.OrderGet.SetPayload(orderDTOs)
	c.EventDispatcher.Dispatch(c.OrderGet)

	return orderDTOs, nil
}

func (c *GetOrdersUseCase) GetByID(id string) (*OrderOutputDTO, error) {
	order, err := c.OrderRepository.GetByID(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("nenhum registro encontrado para o ID: %s", id)
		}
		return nil, err
	}

	orderDTO := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	c.OrderGet.SetPayload(orderDTO)
	c.EventDispatcher.Dispatch(c.OrderGet)

	return &orderDTO, nil
}
