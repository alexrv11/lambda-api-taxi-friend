package services

import (
	"taxifriend/models"
	"taxifriend/repository"
	"taxifriend/utils"
	"github.com/google/uuid"
)

//Driver service
type Driver struct {
	driverRepository repository.IDriver
}

//NewDriver creates a driver service
func NewDriver( driverRepo repository.IDriver) IDriver {
	return &Driver{driverRepository: driverRepo }
}

//Create register a new driver
func (d *Driver) Create(driver *models.Driver) (*models.DriverInfo, error)  {
	driver.Password = utils.CreatePassword(driver.Password)
	driver.ID = uuid.New().String()
	driver.Status = "Registered"
	driver.Credit = 0
	result := &models.DriverInfo{
		ID: driver.ID,
		Status: driver.Status,
		Name: driver.Name,
		CarIdentity: driver.CarIdentity,
		Phone: driver.Phone,
		Credit: driver.Credit,
	}
	return result, d.driverRepository.Create(driver)
}

//GetDriverLocations all driver's location 
func (d *Driver) GetDriverLocations(radio, latitude, longitude float64) ([]models.DriverLocation, error) {
	drivers, err := d.driverRepository.GetDriverLocations()

	if err != nil {
		return nil, err
	}

	result := make([]models.DriverLocation, 0)
	for _, driver := range drivers {
		if radio >= utils.DistanceInKmBetweenEarthCoordinates(latitude, longitude, driver.Latitude, driver.Longitude) {
			item := models.DriverLocation{
				ID:driver.ID,
				Latitude: driver.Latitude,
				Longitude: driver.Longitude,
				Status: driver.Status,
				Direction: driver.Direction,
			}
			result = append(result, item)
		}
	}

	return result, nil
}

//GetItem gets a driver information
func (d *Driver) GetItem(driverID string) (*models.Driver, error)  {
	return d.driverRepository.GetItem(driverID)
}

//UpdateLocation updates a driver's location
func (d *Driver) UpdateLocation(driverID string, location *models.Location) error {

	return d.driverRepository.UpdateLocation(driverID, location)
}