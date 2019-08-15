#!/usr/bin/env bash

#doesnt work, idk why
# protoc -I internal/messages *.proto --go_out=plugins=grpc:pkg/messages

protoc -I configs/protobuf item.proto --go_out=plugins=grpc:./
protoc -I configs/protobuf commerce_listings.proto --go_out=plugins=grpc:./
protoc -I configs/protobuf commerce_prices.proto --go_out=plugins=grpc:./