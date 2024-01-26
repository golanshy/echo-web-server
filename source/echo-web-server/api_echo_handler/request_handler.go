package api_echo_handler

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golanshy/echo-web-server/source/echo-web-server/handler"
	"net/http"
)

func (l EchoLambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (handler.Response, error) {

	res := handler.Response{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Cache-Control":                    "no-cache; no-store",
			"Content-Type":                     "application/json",
		},
	}

	return echo(req, res)
}

func echo(req events.APIGatewayProxyRequest, res handler.Response) (handler.Response, error) {
	lambdaResponse := LambdaResponse{
		ReqMethod:     req.HTTPMethod,
		ReqBody:       req.Body,
		ReqDomainName: req.RequestContext.DomainName,
		ReqHeaders:    req.Headers,
	}
	response, err := json.Marshal(lambdaResponse)

	res.StatusCode = http.StatusOK
	res.Body = string(response)
	return res, err
}
