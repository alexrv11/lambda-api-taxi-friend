package services

import (
	"github.com/alexrv11/lambda-api-taxi-friend/models"
	mockProviders "github.com/alexrv11/lambda-api-taxi-friend/providers/storage/mocks"
	"github.com/alexrv11/lambda-api-taxi-friend/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewDriver(t *testing.T) {
	mockRepository := &mocks.IDriver{}
	mockStorage := &mockProviders.Uploader{}
	service := NewDriver(mockRepository, mockStorage)

	driver := &models.Driver{
		Name:              "driver mock",
		ID:                "10001",
		SideCarPhoto:      "mock side car",
		FrontLicensePhoto: "mock front license",
		BackLicensePhoto:  "mock back license",
		FrontCarPhoto:     "front car photo",
		BackCarPhoto:      "back car photo",
		CarIdentity:       "car identity",
		Credit:            0.0,
		Direction:         1,
		Latitude:          0.0,
		Longitude:         0.0,
		Password:          "mock test",
		Phone:             "123445",
		Status:            "Registered",
	}

	mockRepository.On("Create", mock.Anything).Return(nil).Once()
	mockStorage.On("UploadFile", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything,
	).Return(nil).Once()

	result, err := service.Create(driver)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestDriver_GetDriverLocations(t *testing.T) {
	mockRepository := &mocks.IDriver{}
	service := NewDriver(mockRepository, nil)

	drivers := []models.DriverLocation{
		{
			ID:"10001", Status:"Registered",
			Latitude: -34.634621, Longitude: -58.436132,
		},
		{
			ID:"10003", Status:"Registered",
			Latitude: -34.622863, Longitude: -58.441961,
		},
		{
			ID:"10002", Status:"Registered",
			Latitude: -34.625792, Longitude: -58.441568,
		},
	}

	mockRepository.On("GetDriverLocations").Return(drivers, nil).Once()

	radio := 1.0
	latitude := -34.634162
	longitude := -58.439050


	result, err := service.GetDriverLocations(radio, latitude, longitude)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "10001", result[0].ID)
	assert.Equal(t, "10002", result[1].ID)
}

func TestDriver_GetItem(t *testing.T) {
	mockRepository := &mocks.IDriver{}
	service := NewDriver(mockRepository, nil)
	driverID := "10001"

	driver := &models.Driver{ID:"10001", Name:"Test mock"}

	mockRepository.On("GetItem", driverID).Return(driver, nil).Once()
	result, err := service.GetItem(driverID)

	assert.Nil(t, err)
	assert.Equal(t, driver.Name, result.Name)
	assert.Equal(t, driver.ID, result.ID)
}

func TestDriver_UpdateLocation(t *testing.T) {
	mockRepository := &mocks.IDriver{}
	service := NewDriver(mockRepository, nil)

	driverID := "1001"
	location := &models.Location{Longitude:-20.21, Latitude:10.456}

	mockRepository.On("UpdateLocation", driverID, location).
		Return(nil).Once()

	err := service.UpdateLocation(driverID, location)

	assert.Nil(t, err)
}
