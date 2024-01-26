package api_handler

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golanshy/echo-web-server/source/echo-web-server/handler"
	"log"
	"strings"
)

var (
	echoPath = "echo"
	notFound = "not_found"
)

type LambdaHandler struct {
	echoHandler handler.Handler
}

type LambdaResponse struct {
	Message string
}

// NewLambdaHandler -
func NewLambdaHandler(echoHandler handler.Handler) *LambdaHandler {
	return &LambdaHandler{
		echoHandler: echoHandler,
	}
}

func (l LambdaHandler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (handler.Response, error) {

	log.Printf("HandleRequest Path: %s", req.Path)

	if strings.Contains(req.Path, echoPath) {
		return l.echoHandler.HandleRequest(ctx, req)
	} else {

		res := handler.Response{
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "true",
				"Cache-Control":                    "no-cache; no-store",
				"Content-Type":                     "application/json",
				"Accept":                           "application/json",
				"Content-Security-Policy":          "default-src self",
				"Strict-Transport-Security":        "max-age=31536000; includeSubDomains",
				"X-Content-Type-Options":           "nosniff",
				"X-XSS-Protection":                 "1; mode=block",
				"X-Frame-Options":                  "DENY",
			},
		}
		return res, errors.New(notFound)
	}
}
