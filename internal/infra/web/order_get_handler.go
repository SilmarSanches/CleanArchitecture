package web

import (
	"encoding/json"
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	web "github.com/devfullcycle/20-CleanArch/internal/infra/web/dtos"
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
		web.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		web.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *WebOrderGetHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		web.WriteErrorResponse(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	getOrders := usecase.NewGetOrdersUseCase(h.OrderRepository, h.OrderGetEvent, h.EventDispatcher)
	output, err := getOrders.GetByID(id)
	if err != nil {
		web.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
