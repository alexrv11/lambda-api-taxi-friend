package services

import "taxifriend/models"

//IDriver service interface
type IDriver interface {
	Create(driver *models.Driver) (*models.DriverInfo, error)
	GetDriverLocations(radio, latitude, longitude float64) ([]models.DriverLocation, error)
	GetItem(driverID string) (*models.Driver, error)
	UpdateLocation(driverID string, location *models.Location) error
}
