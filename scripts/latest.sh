#!/bin/sh
# Login docker
echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
# Build Golang Application
go build -o dist/ssh-microservice cmd/main.go
# Build docker image
docker build . -t kainonly/ssh-microservice:latest
# Push docker image
docker push kainonly/ssh-microservice:latest