package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"taxifriend/common/db"
	"taxifriend/common/response"
	"taxifriend/repository"
	"taxifriend/services"
	"taxifriend/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var database = db.CreateDB()
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var handleResponse = response.New(errorLogger)
var qrRepository = repository.NewQr(database)
var qrService = services.NewQr(qrRepository)

//GetQr get order information
func GetQr(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	qrID := req.PathParameters["qrid"]

	result, err := qrService.GetItem(qrID)
	if err != nil {
		return handleResponse.ServerError(err)
	}

	if result == nil {
		return handleResponse.ClientError(http.StatusBadRequest, fmt.Sprintf("Not found %s", qrID))
	}

	data, err := json.Marshal(result)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(data),
	}, nil
}

//CreateQr create an qr
func CreateQr(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload := []byte(req.Body)
	qr := &models.Qr{}
	err := json.Unmarshal(payload, qr)
	
	if err != nil {
		errorLogger.Printf("%s",  err.Error())
	
		return handleResponse.ServerError(err)
	}

	result , err := qrService.Create(qr)
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

//UpdateQrDriverID updates an qr driver
func UpdateQrDriverID(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	qrID := req.PathParameters["qrid"]
	payload := []byte(req.Body)


	qr := &models.InputQrDriver{}
	err := json.Unmarshal(payload, qr)
	
	if err != nil {
		errorLogger.Printf("%s",  err.Error())
	
		return handleResponse.ServerError(err)
	}

	err = qrService.UpdateDriver(qrID, qr.DriverID)
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
			return GetQr(req)
	case "POST":
			return CreateQr(req)
	case "PATCH": 
			return UpdateQrDriverID(req)
	default:
			return handleResponse.ClientError(http.StatusMethodNotAllowed, "resources method not allowed")
	}
}

func main() {
	lambda.Start(router)
}
