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
var orderRepository = repository.NewOrder(database)
var orderService = services.NewOrder(orderRepository)

//GetOrder get order information
func GetOrder(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	orderID := req.PathParameters["orderid"]

	result, err := orderService.GetItem(orderID)
	if err != nil {
		return handleResponse.ServerError(err)
	}

	if result == nil {
		return handleResponse.ClientError(http.StatusBadRequest, fmt.Sprintf("Not found %s", orderID))
	}
	data, err := json.Marshal(result)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(data),
	}, nil
}

//CreateOrder create an order
func CreateOrder(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload := []byte(req.Body)
	order := &models.InputOrder{}
	err := json.Unmarshal(payload, order)
	
	if err != nil {
		errorLogger.Printf("%s",  err.Error())
	
		return handleResponse.ServerError(err)
	}

	result , err := orderService.Create(order)
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

//UpdateOrderStatus updates an order
func UpdateOrderStatus(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	orderID := req.PathParameters["orderid"]
	payload := []byte(req.Body)


	orderStatus := &models.InputOrderStatus{}
	err := json.Unmarshal(payload, orderStatus)
	
	if err != nil {
		errorLogger.Printf("%s",  err.Error())
	
		return handleResponse.ServerError(err)
	}

	err = orderService.UpdateStatus(orderID, orderStatus.Status)
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
			return GetOrder(req)
	case "POST":
			return CreateOrder(req)
	case "PATCH": 
			return UpdateOrderStatus(req)
	default:
			return handleResponse.ClientError(http.StatusMethodNotAllowed, "resources method not allowed")
	}
}

func main() {
	lambda.Start(router)
}
