package services

import (
	"alex/taxi-server/models"
	"alex/taxi-server/repository"
	"github.com/google/uuid"
	"time"
)

type IOrder interface {
	Create(order *models.InputOrder) error
	Update(order *models.InputOrder) error
	GetItem(orderId string) (*models.InputOrder, error)
}

type Order struct {
	OrderRepository repository.IOrder
}



func (o *Order) Create(order *models.InputOrder) error {
	order.ID = uuid.New().String()
	order.Status = "Registered"
	order.DateCreated = time.Now()
	order.LastUpdated = time.Now()
	return o.OrderRepository.Create(order)
}

func (o *Order) Update(order *models.InputOrder) error {

	order.LastUpdated = time.Now()
	return o.OrderRepository.UpdateStatus(order.ID, order.DriverID, order.Status)
}

func (o *Order) GetItem(orderId string) (*models.InputOrder, error) {
	return o.OrderRepository.Get(orderId)
}