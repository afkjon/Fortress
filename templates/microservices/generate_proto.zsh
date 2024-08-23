#!/bin/zsh
protoc --go_out=. --go-grpc_out=. service.proto