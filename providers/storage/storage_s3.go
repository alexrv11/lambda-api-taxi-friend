package storage

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/alexrv11/lambda-api-taxi-friend/providers"
	"github.com/alexrv11/lambda-api-taxi-friend/providers/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

const (
	S3Bucket = "taxi-friend-drivers"
)

type UploaderS3 struct {
	storage providers.Storage
}

func NewUploaderS3(storage providers.Storage) *UploaderS3 {
	return &UploaderS3{storage: storage}
}

func (d *UploaderS3) UploadFile(owner string, contents ...*domain.File) error {

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
