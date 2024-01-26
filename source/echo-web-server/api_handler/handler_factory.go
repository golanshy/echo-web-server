package api_handler

import (
	"github.com/golanshy/echo-web-server/source/echo-web-server/api_echo_handler"
	"github.com/golanshy/echo-web-server/source/echo-web-server/handler"
)

// Create -
func Create() handler.Handler {

	echoHandler := api_echo_handler.Create()

	return NewLambdaHandler(echoHandler)
}
