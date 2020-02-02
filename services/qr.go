package services

import (
	"alex/taxi-server/models"
	"alex/taxi-server/repository"
	"github.com/google/uuid"
)

type IQr interface {
	UpdateDriver(driverId, qrCode string) error
	Create(qr models.Qr) error
	GetItem(code string) (*models.Qr, error)
}

type Qr struct {
	QrRepository repository.IQr
}

func (q *Qr) UpdateDriver(driverID, qrCode string) error  {
	return q.QrRepository.UpdateDriver(driverID, qrCode)
}

func (q *Qr) Create(qr models.Qr) error {
	qr.DriverID = "none"
	qr.Code = uuid.New().String()
	qr.Status = "Registered"
	return q.QrRepository.Create(qr)
}

func (q *Qr) GetItem(code string) (*models.Qr, error) {
	return q.QrRepository.Get(code)
}

