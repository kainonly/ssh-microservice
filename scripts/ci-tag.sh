#!/bin/sh
# Login docker
echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
# Build Golang Application
go build -o dist/ssh-microservice
# Build docker image
docker build . -t kainonly/ssh-microservice:${TRAVIS_TAG}
# Push docker image
docker push kainonly/ssh-microservice:${TRAVIS_TAG}