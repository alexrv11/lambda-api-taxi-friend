package services

import (
	"fmt"
	"github.com/google/uuid"
	"taxifriend/models"
	"taxifriend/providers"
	"taxifriend/providers/domain"
	"taxifriend/repository"
	"taxifriend/utils"
)


//Driver service
type Driver struct {
	driverRepository repository.IDriver
	storage providers.Uploader
}

//NewDriver creates a driver service
func NewDriver( driverRepo repository.IDriver, storage providers.Uploader) IDriver {
	return &Driver{driverRepository: driverRepo, storage: storage}
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

	fileBackCar := &domain.File{
		Name:fmt.Sprintf("%s-taxi-back-photo-%s.jpg",driver.ID, uuid.New().String()),
		Content:driver.BackCarPhoto,
	}

	fileFrontCar := &domain.File{
		Name:fmt.Sprintf("%s-taxi-front-photo-%s.jpg",driver.ID, uuid.New().String()),
		Content:driver.FrontCarPhoto,
	}

	fileBackLicense := &domain.File{
		Name:fmt.Sprintf("%s-license-back-photo-%s.jpg",driver.ID, uuid.New().String()),
		Content: driver.BackLicensePhoto,
	}

	fileFrontLicense := &domain.File{
		Name:fmt.Sprintf("%s-license-front-photo-%s.jpg",driver.ID, uuid.New().String()),
		Content: driver.FrontLicensePhoto,
	}

	fileSideCar := &domain.File{
		Name:fmt.Sprintf("%s-taxi-side-photo-%s.jpg",driver.ID, uuid.New().String()),
		Content: driver.SideCarPhoto,
	}

	err := d.storage.UploadFile(driver.ID, fileFrontLicense, fileBackLicense, fileBackCar, fileFrontCar, fileSideCar)

	if err != nil {
		return nil, err
	}

	driver.FrontCarPhoto = fileFrontCar.Name
	driver.BackCarPhoto = fileBackCar.Name
	driver.SideCarPhoto = fileSideCar.Name
	driver.BackLicensePhoto = fileBackLicense.Name
	driver.FrontLicensePhoto = fileFrontLicense.Name

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