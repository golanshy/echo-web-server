package api_echo_handler

type EchoLambdaHandler struct {
}

// Create -
func Create() *EchoLambdaHandler {
	return NewLambdaHandler()
}

// NewLambdaHandler -
func NewLambdaHandler() *EchoLambdaHandler {
	return &EchoLambdaHandler{}
}
