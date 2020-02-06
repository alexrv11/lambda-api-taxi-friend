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

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var database = db.CreateDB()
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var handleResponse = response.New(errorLogger)
var driverRepositiory = repository.NewDriver(database)

//GetItem gets item
func GetItem(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	driverID := req.PathParameters["driverid"]

	result, err := driverRepositiory.GetItem(driverID)
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

func main() {
	lambda.Start(GetItem)
}
