/*
* API Gateway
*/

resource "aws_apigatewayv2_integration" "lambda_echo_api" {
  api_id = aws_apigatewayv2_api.echo_api.id

  integration_uri    = aws_lambda_function.func.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "get_echo" {
  api_id = aws_apigatewayv2_api.echo_api.id

  route_key = "GET /echo"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_echo_api.id}"
}

resource "aws_apigatewayv2_route" "post_echo" {
  api_id = aws_apigatewayv2_api.echo_api.id

  route_key = "POST /echo"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_echo_api.id}"
}

resource "aws_apigatewayv2_route" "put_echo" {
  api_id = aws_apigatewayv2_api.echo_api.id

  route_key = "PUT /echo"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_echo_api.id}"
}

resource "aws_apigatewayv2_route" "delete_echo" {
  api_id = aws_apigatewayv2_api.echo_api.id

  route_key = "DELETE /echo"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_echo_api.id}"
}

resource "aws_apigatewayv2_route" "patch_echo" {
  api_id = aws_apigatewayv2_api.echo_api.id

  route_key = "PATCH /echo"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_echo_api.id}"
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.func.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.echo_api.execution_arn}/*/*"
}

