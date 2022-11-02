package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"net/http"
	"strings"
)

type Request struct {
	Data string `json:"data"`
}

func parseData(req string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(req)
	if err != nil {
		return "", err
	}
	cmd := strings.Split(string(data), ",")
	return cmd[2], nil
}

func handler(ctx context.Context, request Request) (events.APIGatewayProxyResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	fmt.Println("request data : ", request)
	fmt.Println("aws request ID : ", lc.AwsRequestID)
	iotCode, err := parseData(request.Data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	fmt.Println("iotCode : ", iotCode)

	return events.APIGatewayProxyResponse{
		StatusCode:        http.StatusOK,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              "",
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
