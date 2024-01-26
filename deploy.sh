#!/bin/bash

echo -e "\n+++++ Starting deployment +++++\n"

tfswitch --show-latest
#tfswitch 1.0.0

rm -rf ./bin

echo "+++++ build go packages +++++"

cd source/echo-web-server || exit
go test ./...
env GOOS=linux GOARCH=amd64 go build -o ../../bin/echo-web-server
cd ../..

echo "+++++ hello module +++++"
cd infrastructure || exit
terraform init -input=false
terraform apply -input=false -auto-approve
cd ../

echo -e "\n+++++ Deployment done +++++\n"