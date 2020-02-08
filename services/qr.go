package services

import (
	"taxifriend/models"
	"taxifriend/repository"
	"github.com/google/uuid"
)

//IQr service
type IQr interface {
	UpdateDriver(driverID, qrCode string) error
	Create(qr *models.Qr) (*models.Qr, error)
	GetItem(code string) (*models.Qr, error)
}

//Qr model
type Qr struct {
	QrRepository repository.IQr
}

//NewQr qr service
func NewQr(repo repository.IQr) IQr {
	return &Qr{QrRepository: repo}
}

//UpdateDriver update driver qr
func (q *Qr) UpdateDriver(driverID, qrCode string) error  {
	return q.QrRepository.UpdateDriver(driverID, qrCode)
}

//Create creates qr
func (q *Qr) Create(qr *models.Qr) (*models.Qr, error) {
	code := uuid.New().String()
	qr.DriverID = "none"
	qr.ID = code[:12]
	qr.Status = "Registered"
	return qr, q.QrRepository.Create(qr)
}

//GetItem get item
func (q *Qr) GetItem(code string) (*models.Qr, error) {
	return q.QrRepository.Get(code)
}

