package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"taxifriend/common/db"
	"taxifriend/common/response"
	"taxifriend/models"
	"taxifriend/providers/storage"
	"taxifriend/repository"
	"taxifriend/services"

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

func GetAll(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	radio, err := strconv.ParseFloat(req.QueryStringParameters["radio"], 64)
	latitude, err := strconv.ParseFloat(req.QueryStringParameters["latitude"], 64)
	longitude, err := strconv.ParseFloat(req.QueryStringParameters["longitude"], 64)
	result, err := driverService.GetDriverLocations(radio, latitude, longitude)

	if err != nil {
		return handleResponse.ServerError(err)
	}

	if result == nil {
		msgFormat := fmt.Sprintf("Bad request radio: %f, latitude: %f,longitute: %f", radio, latitude, longitude)
		return handleResponse.ClientError(http.StatusBadRequest, msgFormat)
	}

	wrapRes := &models.Response{Result: result}

	data, err := json.Marshal(wrapRes)

	if err != nil {
		return handleResponse.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(data),
	}, nil
}

func main() {
	lambda.Start(GetAll)
}
