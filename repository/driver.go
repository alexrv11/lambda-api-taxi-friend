package repository

import (
	"fmt"
	"github.com/alexrv11/lambda-api-taxi-friend/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

//IDriver interface
type IDriver interface {
	Create(driver *models.Driver) error
	GetDriverLocations() ([]models.DriverLocation, error)
	GetItem(driverID string) (*models.Driver, error)
	UpdateLocation(driverID string, location *models.Location) error
}

//Driver repository
type Driver struct {
	DB *dynamodb.DynamoDB
}

//NewDriver creates a new driver repository
func NewDriver(db *dynamodb.DynamoDB) IDriver {
	return &Driver{DB: db}
}

//Create creates a new driver entity
func (d *Driver) Create(driver *models.Driver) error {
	tableName := "Driver"
	
	av, err := dynamodbattribute.MarshalMap(driver)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput {
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = d.DB.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

//GetDriverLocations get all driver information
func (d *Driver) GetDriverLocations() ([]models.DriverLocation, error) {

	drivers := make([]models.DriverLocation, 0)
	tableName := "Driver"

	idField := expression.Name("id")
	latitudeField := expression.Name("latitude")
	longitudeField := expression.Name("longitude")
	statusField := expression.Name("status")
	directionField := expression.Name("direction")

	proj := expression.NamesList(idField, latitudeField, longitudeField, statusField, directionField)

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return nil, err
	}
	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := d.DB.Scan(params)
	if err != nil {
		return nil, err
	}

	for _, i := range result.Items {
		item := models.DriverLocation{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return nil, err
		}

		drivers = append(drivers, item)
	}

	return drivers, nil
}

//GetItem gets a driver
func (d *Driver) GetItem(driverID string) (*models.Driver, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Driver"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(driverID),
			},
		},
	}
	result, err := d.DB.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, err
	}

	bk := new(models.Driver)
	err = dynamodbattribute.UnmarshalMap(result.Item, bk)
	if err != nil {
		return nil, err
	}

	return bk, err
}

//UpdateLocation update the driver location
func (d *Driver) UpdateLocation(driverID string, location *models.Location) error {
	
	input := &dynamodb.UpdateItemInput{
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":latitude": {
					N: aws.String(fmt.Sprintf("%f", location.Latitude)),
				},
				":longitude": {
					N: aws.String(fmt.Sprintf("%f", location.Longitude)),
				},
		},
    TableName: aws.String("Driver"),
    Key: map[string]*dynamodb.AttributeValue{
        "id": {
            S: aws.String(driverID),
        },
    },
    ReturnValues:     aws.String("UPDATED_NEW"),
    UpdateExpression: aws.String("set latitude = :latitude, longitude = :longitude"),
	}

	_, err := d.DB.UpdateItem(input)
	if err != nil {
    return err
	}
	
	return nil
}
