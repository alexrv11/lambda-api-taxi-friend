package repository

import (
	"taxifriend/models"
)

//IOrder repository interface
type IOrder interface {
	Create(order *models.InputOrder) error
	UpdateStatus(id, driverID, status string) error
	Get(id string) (*models.InputOrder, error)
}


//Order repository implementation
type Order struct {
	//db  *dynamo.DB
}

//Get order by id
func (o *Order) Get(id string) (*models.InputOrder, error)  {
	/*table := o.db.Table("Order")
	var result models.InputOrder

	err := table.Get("Id", id).One(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil*/
	return nil, nil
}

//UpdateStatus updates an order by driverId and status
func (o *Order) UpdateStatus(id, driverID, status string) error {
	/*table := o.db.Table("Order")

	err := table.Update("Id", id).
		If("DriverId = ?", driverId).
		Set("Status", status).
		Set("LastUpdated", time.Now()).
		Run()

	return err*/
	return nil
}

//Create creates an order
func (o *Order) Create(order *models.InputOrder) error {
	/*table := o.db.Table("Order")

	err := table.Put(order).Run()
	*/
	return nil
}
