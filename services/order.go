package services

import (
	"github.com/alexrv11/lambda-api-taxi-friend/common/util"
	"github.com/alexrv11/lambda-api-taxi-friend/models"
	"github.com/alexrv11/lambda-api-taxi-friend/repository"
	"time"

	"github.com/google/uuid"
)

//IOrder interface
type IOrder interface {
	Create(inputOrder *models.InputOrder) (*models.Order, error)
	UpdateStatus(id, status string) error
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
func (o *Order) Create(inputOrder *models.InputOrder) (*models.Order, error) {

	id := uuid.New().String()
	status := "Registered"
	dateCreated := time.Now().UTC().Format(util.FormatDate)
	lastUpdated := time.Now().UTC().Format(util.FormatDate)

	order := &models.Order{
		ID:          id,
		Status:      status,
		DateCreated: dateCreated,
		LastUpdated: lastUpdated,
		Longitude:   inputOrder.Longitude,
		Latitude:    inputOrder.Latitude,
		DriverID:    inputOrder.DriverID,
	}

	return order, o.OrderRepository.Create(order)
}

//UpdateStatus updates a order
func (o *Order) UpdateStatus(id, status string) error {

	return o.OrderRepository.UpdateStatus(id, status)
}

//GetItem gets an item
func (o *Order) GetItem(orderID string) (*models.Order, error) {
	return o.OrderRepository.Get(orderID)
}
