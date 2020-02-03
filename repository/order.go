package repository

import (
	"taxifriend/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//IOrder repository interface
type IOrder interface {
	Create(order *models.Order) error
	UpdateStatus(id, driverID, status string) error
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
	/*table := o.db.Table("Order")
	var result models.InputOrder

	err := table.Get("Id", id).One(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil*/
	return nil, nil
}

//UpdateStatus updates an order by driverId and status
func (o *Order) UpdateStatus(id, driverID, status string) error {
	/*table := o.db.Table("Order")

	err := table.Update("Id", id).
		If("DriverId = ?", driverId).
		Set("Status", status).
		Set("LastUpdated", time.Now()).
		Run()

	return err*/
	return nil
}

//Create creates an order
func (o *Order) Create(order *models.Order) error {
	tableName := "Order"
	av, err := dynamodbattribute.MarshalMap(order)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput {
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = o.DB.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
