#!/bin/bash

# Kill any connection for postgresql container to connect
sudo fuser -n tcp -k 5432

# Set up postgresql container
#go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
CGO_ENABLED=0 go build -a -installsuffix cgo
sudo docker-compose up -d --build

sleep 5 && sudo docker-compose logs api && sudo docker-compose ps
