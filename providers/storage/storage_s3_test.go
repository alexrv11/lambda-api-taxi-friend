package storage

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strings"
	"github.com/alexrv11/lambda-api-taxi-friend/providers/domain"
	"github.com/alexrv11/lambda-api-taxi-friend/providers/storage/mocks"
	"testing"
)

func TestUploaderS3_UploadFile(t *testing.T) {
	mockStorage := &mocks.Storage{}
	mockStorage.On("PutObject", mock.MatchedBy(func(input *s3.PutObjectInput) bool {
		key := *input.Key
		if strings.EqualFold(key, "1000/test1") ||
			 strings.EqualFold(key ,"1000/test2") {
			return true
		}

		return false
	})).Return(nil, nil).Twice()

	uploader := &UploaderS3{mockStorage}
	file1 := &domain.File{Name: "test1", Content: "Y29udGVudC1tb2Nr"}
	file2 := &domain.File{Name: "test2", Content: "Y29udGVudC1tb2Nr"}
	ownerID := "1000"

	err := uploader.UploadFile(ownerID, file1, file2)

	assert.Nil(t, err)
	mockStorage.AssertExpectations(t)
}
