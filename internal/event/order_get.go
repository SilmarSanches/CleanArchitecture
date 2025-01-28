package event

import "time"

type OrderGet struct {
	Name    string
	Payload interface{}
}

func NewOrderGet() *OrderGet {
	return &OrderGet{
		Name: "OrderGet",
	}
}

func (e *OrderGet) GetName() string {
	return e.Name
}

func (e *OrderGet) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderGet) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderGet) GetDateTime() time.Time {
	return time.Now()
}
