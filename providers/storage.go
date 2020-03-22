package providers

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/alexrv11/lambda-api-taxi-friend/providers/domain"
)

type Storage interface {
	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

type Uploader interface {
	UploadFile(owner string, contents ...*domain.File) error
}