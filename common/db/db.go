package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/logger"
)

// CreateDB creates a DB management
func CreateDB() *dynamodb.DynamoDB {
	var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
	return db
}

//CreateS3 Create S3 management
func CreateS3() (*s3.S3, error) {
	s, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		logger.Errorf("Error creating the s3 service %v", err)
		return nil, err
	}

	return s3.New(s), nil
}
