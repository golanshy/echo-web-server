package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golanshy/echo-web-server/source/echo-web-server/api_handler"
	"log"
)

func main() {
	log.Printf("echo-web-server starting")
	handler := api_handler.Create()
	lambda.Start(handler.HandleRequest)
}
