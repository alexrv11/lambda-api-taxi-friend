package repository

import (
	"github.com/alexrv11/lambda-api-taxi-friend/common/util"
	"github.com/alexrv11/lambda-api-taxi-friend/models"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//IOrder repository interface
type IOrder interface {
	Create(order *models.Order) error
	UpdateStatus(id, status string) error
	Get(id string) (*models.Order, error)
}

//Order repository implementation
type Order struct {
	DB *dynamodb.DynamoDB
}

//NewOrder create a new instance of Order service
func NewOrder(db *dynamodb.DynamoDB) IOrder {
	return &Order{DB: db}
}

//Get order by id
func (o *Order) Get(id string) (*models.Order, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Order"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}
	result, err := o.DB.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, err
	}

	bk := new(models.Order)
	err = dynamodbattribute.UnmarshalMap(result.Item, bk)
	if err != nil {
		return nil, err
	}

	return bk, err
}

//UpdateStatus updates an order by driverId and status
func (o *Order) UpdateStatus(id, status string) error {
	statusName := "status"
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":status": {
				S: aws.String(status),
			},
			":last_updated": {
				S: aws.String(time.Now().UTC().Format(util.FormatDate)),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#ts": &statusName,
		},
		TableName: aws.String("Order"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #ts = :status, last_updated= :last_updated"),
	}

	_, err := o.DB.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}

//Create creates an order
func (o *Order) Create(order *models.Order) error {
	tableName := "Order"
	av, err := dynamodbattribute.MarshalMap(order)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = o.DB.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
