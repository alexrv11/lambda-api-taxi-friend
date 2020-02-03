package main

import (
	"encoding/json"
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

//CreateOrder gets item
func CreateOrder(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload := []byte(req.Body)
	order := &models.InputOrder{}
	err := json.Unmarshal(payload, order)
	
	if err != nil {
		errorLogger.Printf("%s",  err.Error())
	
		return handleResponse.ServerError(err)
	}

	err = orderService.Create(order)
	if err != nil {
		return handleResponse.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "",
	}, nil
}

func main() {
	lambda.Start(CreateOrder)
}
