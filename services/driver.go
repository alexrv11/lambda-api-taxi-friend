package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"taxifriend/models"
	"taxifriend/repository"
	"taxifriend/utils"
	"net/http"
	"github.com/google/uuid"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const(
	S3Bucket="taxi-friend-drivers"
)

type File struct {
	Name string
	Content string
}

//Driver service
type Driver struct {
	driverRepository repository.IDriver
	storage *s3.S3
}

//NewDriver creates a driver service
func NewDriver( driverRepo repository.IDriver, storage *s3.S3) IDriver {
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

	fileBackCar := &File{
		Name:fmt.Sprintf("%s-taxi-back-photo-%s",driver.ID, uuid.New().String()),
		Content:driver.BackCarPhoto,
	}

	fileFrontCar := &File{
		Name:fmt.Sprintf("%s-taxi-front-photo-%s",driver.ID, uuid.New().String()),
		Content:driver.FrontCarPhoto,
	}

	fileBackLicense := &File{
		Name:fmt.Sprintf("%s-license-back-photo-%s",driver.ID, uuid.New().String()),
		Content: driver.BackLicensePhoto,
	}

	fileFrontLicense := &File{
		Name:fmt.Sprintf("%s-license-front-photo-%s",driver.ID, uuid.New().String()),
		Content: driver.FrontLicensePhoto,
	}

	fileSideCar := &File{
		Name:fmt.Sprintf("%s-taxi-side-photo-%s",driver.ID, uuid.New().String()),
		Content: driver.SideCarPhoto,
	}

	err := d.uploadContent(driver.ID, fileFrontLicense, fileBackLicense, fileBackCar, fileFrontCar, fileSideCar)

	return result, d.driverRepository.Create(driver)
}

func (d *Driver) uploadContent(owner string, contents ...*File) error {

	for _, file := range contents {
			buffer, err := base64.StdEncoding.DecodeString(file.Content)
			if err != nil {
				return err
			}
			_, err = d.storage.PutObject(&s3.PutObjectInput{
				Bucket:               aws.String(S3Bucket),
				Key:                  aws.String(fmt.Sprintf("%s/%s", owner, file.Name)),
				ACL:                  aws.String("private"),
				Body:                 bytes.NewReader(buffer),
				ContentLength:        aws.Int64(int64(len(buffer))),
				ContentType:          aws.String(http.DetectContentType(buffer)),
				ContentDisposition:   aws.String("attachment"),
				ServerSideEncryption: aws.String("AES256"),
		})

		if err != nil {
			return err
		}
	}

	return nil
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