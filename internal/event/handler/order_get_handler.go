package handler

import (
	"encoding/json"
	"fmt"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/streadway/amqp"
	"sync"
)

type OrderGetHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderGetHandler(rabbitMQChannel *amqp.Channel) *OrderGetHandler {
	return &OrderGetHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderGetHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order get all: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
