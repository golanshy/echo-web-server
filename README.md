# Start a new Go project

`cd ~/Workspace/projects`

`mkdir echo-web-server`

`cd echo-web-server`

`go mod init github.com/golanshy/echo-web-server`

`mkdir -p cmd/echo-web-server`

`touch cmd/echo-web-server/main.go`


# echo-web-server setup

Setup a Go Lambda function with Terraform.

## Setup

`chmod +x ./deploy.sh`

`chmod +x ./destroy.sh`

## Usage

- To deploy run `sh deploy.sh`

- To destroy run `sh destroy.sh`
