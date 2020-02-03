package services

import (
	"taxifriend/models"
	"taxifriend/repository"
	"time"

	"github.com/google/uuid"
)

//IOrder interface
type IOrder interface {
	Create(inputOrder *models.InputOrder) error
	Update(order *models.Order) error
	GetItem(orderID string) (*models.Order, error)
}

//Order service
type Order struct {
	OrderRepository repository.IOrder
}

//NewOrder create a order service
func NewOrder(orderRepository repository.IOrder) IOrder {
	return &Order{OrderRepository: orderRepository}
}

//Create register a new order
func (o *Order) Create(inputOrder *models.InputOrder) error {

	id := uuid.New().String()
	status := "Registered"
	dateCreated := time.Now()
	lastUpdated := time.Now()

	order := &models.Order {
						ID: id,
						Status : status,
						DateCreated: dateCreated,
						LastUpdated:lastUpdated,
						Longitude: inputOrder.Longitude,
						Latitude: inputOrder.Latitude,
						DriverID: inputOrder.DriverID,
					}
	
	return o.OrderRepository.Create(order)
}

//Update updates a order
func (o *Order) Update(order *models.Order) error {
	order.LastUpdated = time.Now()
	return o.OrderRepository.UpdateStatus(order.ID, order.DriverID, order.Status)
}

//GetItem gets an item
func (o *Order) GetItem(orderID string) (*models.Order, error) {
	return o.OrderRepository.Get(orderID)
}
