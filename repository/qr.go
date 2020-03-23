package repository

import (
	"github.com/alexrv11/lambda-api-taxi-friend/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//IQr repository interface
type IQr interface {
	Get(code string) (*models.Qr, error)
	UpdateDriver(driverID, qr string) error
	Create(qr *models.Qr) error
}

//Qr repository
type Qr struct {
	DB *dynamodb.DynamoDB
}

//NewQr new qr
func NewQr(db *dynamodb.DynamoDB) IQr {
	return &Qr{DB: db}
}

//Get qr information by id
func (q *Qr) Get(id string) (*models.Qr, error) {

	input := &dynamodb.GetItemInput{
		TableName: aws.String("CreditQr"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := q.DB.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, err
	}

	bk := new(models.Qr)
	err = dynamodbattribute.UnmarshalMap(result.Item, bk)
	if err != nil {
		return nil, err
	}

	return bk, err
}

//UpdateDriver updates the qr's driver information
func (q *Qr) UpdateDriver(driverID, qr string) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":driver_id": {
				S: aws.String(driverID),
			},
		},
		TableName: aws.String("CreditQr"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(qr),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set driver_id = :driver_id"),
	}

	_, err := q.DB.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}

//Create creates a new qr
func (q *Qr) Create(qr *models.Qr) error {
	tableName := "CreditQr"
	av, err := dynamodbattribute.MarshalMap(&qr)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = q.DB.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
