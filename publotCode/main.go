package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
)

type PayloadReq struct {
	Command string `json:"command"`
}

type PubInput struct {
	Payload []byte
	Qos     int32
	Retain  bool
	Topic   string
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) error {
	lc, _ := lambdacontext.FromContext(ctx)
	fmt.Println(lc)

	fmt.Println("request.Body :", request.Body)

	cmdReq := PayloadReq{}

	err := json.Unmarshal([]byte(request.Body), &cmdReq)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
	}

	svc := iotdataplane.NewFromConfig(cfg)

	pub := PubInput{
		Payload: []byte(cmdReq.Command),
		Qos:     1,
		Retain:  true,
		Topic:   "iot/test/pub",
	}

	input := iotdataplane.PublishInput{
		Payload: pub.Payload,
		Qos:     pub.Qos,
		Retain:  pub.Retain,
		Topic:   &pub.Topic,
	}

	output, err := svc.Publish(context.TODO(), &input)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("output : ", output)

	return nil
}

func main() {
	lambda.Start(handler)
}
