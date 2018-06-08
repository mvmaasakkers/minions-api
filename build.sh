#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o ./build/api cmd/main.go

docker build -t eu.gcr.io/bb-hackathon/hackathon-api:latest .

docker push eu.gcr.io/bb-hackathon/hackathon-api:latest

# kubectl set image deployment/api api=eu.gcr.io/bb-hackathon/hackathon-api:latest