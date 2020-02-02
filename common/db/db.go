package db


import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
)

// CreateDB creates a DB management
func CreateDB() *dynamodb.DynamoDB {
	var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
	return db
}