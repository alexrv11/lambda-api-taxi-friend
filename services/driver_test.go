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
