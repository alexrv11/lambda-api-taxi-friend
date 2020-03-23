package response

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

//HandlerResponse prepare a gateway response
type HandlerResponse struct {
	log *log.Logger
}

//New creates a *HandlerResponse instance
func New(log *log.Logger) *HandlerResponse {
	return &HandlerResponse{log: log}
}

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func (handle *HandlerResponse) ServerError(err error) (events.APIGatewayProxyResponse, error) {
	handle.log.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
	}, nil
}

// Similarly add a helper for send responses relating to client errors.
func (handle *HandlerResponse) ClientError(status int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
