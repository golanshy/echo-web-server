/*
* Cloudwatch
*/

// Log group

resource "aws_cloudwatch_log_group" "echo_api_gw" {
  name = "/aws/lambda/echo-api"
  retention_in_days = 14
}
