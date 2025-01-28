package web

import (
	"encoding/json"
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"net/http"
)

type WebOrderGetHandler struct {
	EventDispatcher events.EventDispatcherInterface
	OrderRepository entity.OrderRepositoryInterface
	OrderGetEvent   events.EventInterface
}

func NewWebOrderGetHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderGetEvent events.EventInterface,
) *WebOrderGetHandler {
	return &WebOrderGetHandler{
		EventDispatcher: EventDispatcher,
		OrderRepository: OrderRepository,
		OrderGetEvent:   OrderGetEvent,
	}
}

func (h *WebOrderGetHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	getOrders := usecase.NewGetOrdersUseCase(h.OrderRepository, h.OrderGetEvent, h.EventDispatcher)
	output, err := getOrders.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
