package main

import (
        "fmt"
        "context"
    	"encoding/json"
        "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

type MyResponse struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context) (events.APIGatewayProxyResponse, error) {
        msgObject := getMsg();
	serializedMsg, _ := json.Marshal(msgObject)
	return events.APIGatewayProxyResponse{Body: string(serializedMsg), StatusCode: 200}, nil
}

func getMsg() (MyResponse) {
	return MyResponse{fmt.Sprintf("Hello, %s!", "main-only, no get build from bin" )}
}

func main() {
        lambda.Start(HandleRequest)
}
