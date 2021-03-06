package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexrv11/lambda-api-taxi-friend/common/db"
	"github.com/alexrv11/lambda-api-taxi-friend/common/response"
	"github.com/alexrv11/lambda-api-taxi-friend/models"
	"github.com/alexrv11/lambda-api-taxi-friend/providers/storage"
	"github.com/alexrv11/lambda-api-taxi-friend/repository"
	"github.com/alexrv11/lambda-api-taxi-friend/services"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var database = db.CreateDB()
var s3, _ = db.CreateS3()
var uploader = storage.NewUploaderS3(s3)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var handleResponse = response.New(errorLogger)
var driverRepository = repository.NewDriver(database)
var driverService = services.NewDriver(driverRepository, uploader)

//GetItem gets item
func GetItem(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	driverID := req.PathParameters["driverid"]

	result, err := driverService.GetItem(driverID)
	if err != nil {
		return handleResponse.ServerError(err)
	}

	if result == nil {
		return handleResponse.ClientError(http.StatusBadRequest, fmt.Sprintf("Not found %s", driverID))
	}
	data, err := json.Marshal(result)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(data),
	}, nil
}

//Create create an driver
func Create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload := []byte(req.Body)
	driver := &models.Driver{}
	err := json.Unmarshal(payload, driver)

	if err != nil {
		errorLogger.Printf("%s", err.Error())

		return handleResponse.ServerError(err)
	}

	result, err := driverService.Create(driver)
	if err != nil {
		return handleResponse.ServerError(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		return handleResponse.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(data),
	}, nil
}

//UpdateDriverLocation updates an driver location
func UpdateDriverLocation(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	driverID := req.PathParameters["driverid"]
	payload := []byte(req.Body)

	location := &models.Location{}
	err := json.Unmarshal(payload, location)

	if err != nil {
		errorLogger.Printf("%s", err.Error())

		return handleResponse.ServerError(err)
	}

	err = driverService.UpdateLocation(driverID, location)
	if err != nil {
		return handleResponse.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "",
	}, nil
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return GetItem(req)
	case "POST":
		return Create(req)
	case "PATCH":
		return UpdateDriverLocation(req)
	default:
		return handleResponse.ClientError(http.StatusMethodNotAllowed, "resources method not allowed")
	}
}

func main() {
	lambda.Start(router)
}
