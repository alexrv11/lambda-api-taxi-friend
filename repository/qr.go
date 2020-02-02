package repository

import (
	"taxifriend/models"
)

//IQr repository interface
type IQr interface {
	Get(code string) (*models.Qr, error)
	UpdateDriver(driverID, qr string)  error
	Create(qr models.Qr) error
}

//Qr repository
type Qr struct {
	//db  *dynamo.DB
}

//Get qr information by code
func (q *Qr) Get(code string) (*models.Qr, error)  {
	/*table := q.db.Table("CreditQr")
	var result models.Qr

	err := table.Get("Code", code).One(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil*/
	return nil, nil
}

//UpdateDriver updates the qr's driver information
func (q *Qr) UpdateDriver(driverID, qr string) error {
	/*table := q.db.Table("CreditQr")

	err := table.Update("Code", qr).
		If("DriverID = ?", "none").
		Set("DriverID", driverID).
		Run()

	return err*/

	return nil
}

//Create creates a new qr
func (q *Qr) Create(qr models.Qr) error {
	/*table := q.db.Table("CreditQr")

	err := table.Put(qr).Run()

	return err*/
	return nil
}